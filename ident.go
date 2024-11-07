package stalker

import (
	"go/ast"
)

func IsIdent(e ast.Expr) bool {
	_, ok := e.(*ast.Ident)
	return ok
}

type Ident struct {
	Name *ast.Ident
	Type string
}

func ParseIdent(names []*ast.Ident, ident *ast.Ident) Ident {
	return Ident{
		Name: names[0],
		Type: ident.Name,
	}
}
