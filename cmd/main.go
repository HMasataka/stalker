package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/HMasataka/stalker"
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

	functions := stalker.ListFunctions(f)

	for _, fn := range functions {
		f := stalker.ParseFn(fn)

		for _, param := range f.Params {
			if !stalker.IsSelectorExpr(param.Type) && !stalker.IsIdent(param.Type) {
				continue
			}

			if stalker.IsSelectorExpr(param.Type) {
				println("selector")
				selector := param.Type.(*ast.SelectorExpr)

				sel := stalker.ParseSelector(param.Names, selector)
				fmt.Printf("%s %s\n", sel.Name, sel.FullTypeName)
			}

			if stalker.IsIdent(param.Type) {
				println("ident")
				ident := param.Type.(*ast.Ident)

				idt := stalker.ParseIdent(param.Names, ident)
				fmt.Printf("%s %s\n", idt.Name, idt.Type)
			}
		}
	}
}
