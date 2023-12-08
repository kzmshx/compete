package analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type FunctionInfo struct {
	Name string
	Decl *ast.FuncDecl
	Docs *ast.CommentGroup
}

func NewFunctionInfo(decl *ast.FuncDecl) *FunctionInfo {
	return &FunctionInfo{
		Name: decl.Name.Name,
		Decl: decl,
		Docs: decl.Doc,
	}
}

func FindUnusedFunctions(code string) ([]*FunctionInfo, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, 0)
	if err != nil {
		return nil, err
	}

	var unusedFunctions []*FunctionInfo

	usedFunctions := make(map[string]bool)
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			if fn, ok := x.Fun.(*ast.Ident); ok {
				usedFunctions[fn.Name] = true
			}
		}
		return true
	})

	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if fn.Name.Name != "main" && !usedFunctions[fn.Name.Name] {
				unusedFunctions = append(unusedFunctions, NewFunctionInfo(fn))
			}
		}
	}

	return unusedFunctions, nil
}
