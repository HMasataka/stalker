package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `
package main
func A(ctx context.Context, id int, name string) {}
`
	fs := token.NewFileSet()

	f, err := parser.ParseFile(fs, "", src, 0)
	if err != nil {
		panic(err)
	}

	for _, decl := range f.Decls {
		ast.Inspect(decl, func(node ast.Node) bool {
			fn, ok := node.(*ast.FuncDecl)
			if !ok {
				fmt.Println(false)
				return false
			}

			fmt.Println("func name : ", fn.Name.Name)

			for _, param := range fn.Type.Params.List {
				selector, isSelectorExpr := param.Type.(*ast.SelectorExpr)
				ident, isIdent := param.Type.(*ast.Ident)

				if !isSelectorExpr && !isIdent {
					continue
				}

				if isSelectorExpr {
					println("sel")
					typeName := selector.Sel.Name
					typ := selector.X.(*ast.Ident)
					for _, name := range param.Names {
						fmt.Printf("%s %s.%s\n", name.Name, typ, typeName)
					}
				}

				if isIdent {
					println("ident")
					typeName := ident.Name
					for _, name := range param.Names {
						println(name.Name, typeName)
					}
				}
			}

			return false
		})
	}
}
