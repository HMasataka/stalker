package stalker

import (
	"fmt"
	"go/ast"
)

func IsSelectorExpr(e ast.Expr) bool {
	_, ok := e.(*ast.SelectorExpr)
	return ok
}

type Selector struct {
	Name         *ast.Ident
	Type         string
	FullTypeName string
}

func ParseSelector(names []*ast.Ident, selector *ast.SelectorExpr) Selector {
	typeName := selector.Sel.Name
	pkg := selector.X.(*ast.Ident)

	return Selector{
		Name:         names[0],
		Type:         selector.Sel.Name,
		FullTypeName: fmt.Sprintf("%s.%s", pkg, typeName),
	}
}
