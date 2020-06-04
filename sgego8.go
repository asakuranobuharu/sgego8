package sgego8

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "myanalyzer is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "myanalyzer",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.CallExpr)(nil),
		(*ast.AssignStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			if n.Name.Name == "gopher" {
				pass.Reportf(n.Pos(), "function is gopher")
			}
		case *ast.CallExpr:
			switch f := n.Fun.(type) {
				case *ast.Ident:
					if f.Name == "gopher" {
						pass.Reportf(f.Pos(), "call gopher function")
					}
			}
		case *ast.AssignStmt:
			if len(n.Lhs) != 1 {
				return
			}
			switch e := n.Rhs[0].(type) {
			case *ast.TypeAssertExpr:
				pass.Reportf(e.Pos(), "required two values")
			}
		}
	})

	return nil, nil
}
