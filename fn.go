package stalker

import "go/ast"

type Fn struct {
	Name   string
	Params []*ast.Field
}

func ParseFn(f *ast.FuncDecl) Fn {
	return Fn{
		Name:   f.Name.Name,
		Params: f.Type.Params.List,
	}
}

func ListFunctions(f *ast.File) []*ast.FuncDecl {
	var functions []*ast.FuncDecl

	for _, decl := range f.Decls {
		ast.Inspect(decl, func(node ast.Node) bool {
			fn, ok := node.(*ast.FuncDecl)
			if !ok {
				return false
			}

			functions = append(functions, fn)

			return false
		})
	}

	return functions
}
