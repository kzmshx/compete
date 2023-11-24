package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"

	"golang.org/x/tools/go/packages"
	"honnef.co/go/tools/go/loader"
	"honnef.co/go/tools/lintcmd/cache"
	"honnef.co/go/tools/unused"
)

func listUnusedDecls(filename string) ([]string, error) {
	opts := unused.DefaultOptions
	c, err := cache.Default()
	if err != nil {
		return nil, err
	}
	cfg := &packages.Config{}
	specs, err := loader.Graph(c, cfg, filename)
	if err != nil {
		return nil, err
	}

	var sg unused.SerializedGraph
	exists := false
	for _, spec := range specs {
		if len(spec.Errors) != 0 {
			continue
		}

		lpkg, _, err := loader.Load(spec)
		if err != nil {
			continue
		}
		if len(lpkg.Errors) != 0 {
			continue
		}

		g := unused.Graph(lpkg.Fset, lpkg.Syntax, lpkg.Types, lpkg.TypesInfo, nil, nil, opts)
		if len(g) > 0 {
			sg.Merge(g)
			exists = true
		}
	}
	if !exists {
		return nil, nil
	}

	var results []string
	for _, obj := range sg.Results().Unused {
		results = append(results, obj.Name)
	}

	return results, nil
}

func parseFile(filename string) (*ast.File, *token.FileSet, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	return file, fset, err
}

func removeDecl(file *ast.File, name string) {
	// Remove declaration.
	var newDecls []ast.Decl
	removedCommentGroupPos := map[token.Pos]struct{}{}
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Name.Name == name {
				removedCommentGroupPos[d.Doc.Pos()] = struct{}{}
				continue
			}
		case *ast.GenDecl:
			var newSpecs []ast.Spec
			for _, spec := range d.Specs {
				switch s := spec.(type) {
				case *ast.ValueSpec:
					includes := false
					for _, n := range s.Names {
						if n.Name == name {
							includes = true
						}
					}
					if includes {
						removedCommentGroupPos[d.Doc.Pos()] = struct{}{}
						continue
					}
				case *ast.TypeSpec:
					if s.Name.Name == name {
						removedCommentGroupPos[d.Doc.Pos()] = struct{}{}
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
		newDecls = append(newDecls, decl)
	}
	file.Decls = newDecls

	// Remove comments that were attached to the removed declaration.
	var newComments []*ast.CommentGroup
	for _, commentGroup := range file.Comments {
		if _, ok := removedCommentGroupPos[commentGroup.Pos()]; ok {
			continue
		}
		newComments = append(newComments, commentGroup)
	}
	file.Comments = newComments
}

func writeFile(file *ast.File, fset *token.FileSet, filename string) error {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, file); err != nil {
		return err
	}
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

func Run(filename string) error {
	decls, err := listUnusedDecls(filename)
	if err != nil {
		return err
	}
	if len(decls) == 0 {
		return nil
	}

	file, fset, err := parseFile(filename)
	if err != nil {
		return err
	}
	for _, decl := range decls {
		removeDecl(file, decl)
	}
	if err := writeFile(file, fset, filename); err != nil {
		return err
	}
	return nil
}
