package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"golang.org/x/tools/go/packages"
	"honnef.co/go/tools/go/loader"
	"honnef.co/go/tools/lintcmd/cache"
	"honnef.co/go/tools/unused"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s </path/to/main.go>\n", os.Args[0])
		os.Exit(1)
	}

	targetPath := os.Args[1]
	targetDir := filepath.Dir(targetPath)
	if err := os.Chdir(targetDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	mainPath := "./main.go"

	file, fset, err := ParseFile(mainPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := RemoveComments(file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := RemoveUnusedFunctions(file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := RemoveUnusedDeclarations(mainPath, file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := WriteFile(mainPath, file, fset); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ParseFile(filename string) (*ast.File, *token.FileSet, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	return file, fset, err
}

func WriteFile(filename string, file *ast.File, fset *token.FileSet) error {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, file); err != nil {
		return err
	}
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

func RemoveComments(file *ast.File) error {
	file.Doc = nil
	file.Comments = []*ast.CommentGroup{}
	return nil
}

type Pkg struct {
	Name  string
	Path  string
	Funcs []Function
}

type Function struct {
	Name      string
	Position  Position
	Generated bool
}

type Position struct {
	File      string
	Line, Col int
}

type UnusedFunction struct {
	Type string
	Name string
}

func RemoveUnusedFunctions(file *ast.File) error {
	cmd := exec.Command("deadcode", "-json", ".")
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	var pkgs []Pkg
	if err := json.Unmarshal(output, &pkgs); err != nil {
		return err
	}
	if len(pkgs) == 0 {
		return nil
	}
	if len(pkgs) > 1 || pkgs[0].Name != "main" {
		return fmt.Errorf("only single `main` package is supported")
	}

	var unusedFunctions []UnusedFunction
	for _, fn := range pkgs[0].Funcs {
		if names := strings.Split(fn.Name, "."); len(names) > 1 {
			unusedFunctions = append(unusedFunctions, UnusedFunction{Type: names[0], Name: names[1]})
		} else {
			unusedFunctions = append(unusedFunctions, UnusedFunction{Type: "", Name: names[0]})
		}
	}

	var decls []ast.Decl
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if slices.ContainsFunc(unusedFunctions, func(fn UnusedFunction) bool { return fn.Name == d.Name.Name }) {
				continue
			}
		}
		decls = append(decls, decl)
	}
	file.Decls = decls

	return nil
}

func RemoveUnusedDeclarations(filename string, file *ast.File) error {
	unused.Debug = nil

	opts := unused.DefaultOptions

	c, err := cache.Default()
	if err != nil {
		return err
	}

	cfg := &packages.Config{}
	specs, err := loader.Graph(c, cfg, filename)
	if err != nil {
		return err
	}

	var sg unused.SerializedGraph
	exists := false
	for _, spec := range specs {
		if len(spec.Errors) != 0 {
			return fmt.Errorf("errors in %s: %s", spec.PkgPath, spec.Errors)
		}

		lpkg, _, err := loader.Load(spec)
		if err != nil {
			return err
		}
		if len(lpkg.Errors) != 0 {
			return fmt.Errorf("errors in %s: %s", spec.PkgPath, lpkg.Errors)
		}

		g := unused.Graph(lpkg.Fset, lpkg.Syntax, lpkg.Types, lpkg.TypesInfo, nil, nil, opts)
		if len(g) > 0 {
			sg.Merge(g)
			exists = true
		}
	}

	var results []string
	if !exists {
		return nil
	}
	for _, obj := range sg.Results().Unused {
		results = append(results, obj.Name)
	}

	var decls []ast.Decl
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if slices.Contains(results, d.Name.Name) {
				continue
			}
		case *ast.GenDecl:
			var newSpecs []ast.Spec
			for _, spec := range d.Specs {
				switch s := spec.(type) {
				case *ast.ValueSpec:
					includes := false
					for _, n := range s.Names {
						if slices.Contains(results, n.Name) {
							includes = true
						}
					}
					if includes {
						continue
					}
				case *ast.TypeSpec:
					if slices.Contains(results, s.Name.Name) {
						continue
					}
				}
				newSpecs = append(newSpecs, spec)
			}
			if len(newSpecs) == 0 {
				continue
			}
			d.Specs = newSpecs
		}
		decls = append(decls, decl)
	}
	file.Decls = decls

	return nil
}
