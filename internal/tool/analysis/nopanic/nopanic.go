// Package nopanic implements an analyzer that checks if somewhere in the
// source, there is a panic.
package nopanic

import (
	"go/ast"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/ctrlflow"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/cfg"
)

// Analyzer implements the analyzer that checks for panics.
var Analyzer = &analysis.Analyzer{
	Name: "nopanic",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		ctrlflow.Analyzer,
		inspect.Analyzer,
	},
}

// Doc is the documentation string that is shown on the command line if help is
// requested.
const Doc = "check if there is any panic in the code"

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	cf := pass.ResultOf[ctrlflow.Analyzer].(*ctrlflow.CFGs)

	var blocks []*cfg.Block
	inspect.Preorder(
		[]ast.Node{
			(*ast.FuncDecl)(nil),
		},
		func(n ast.Node) {
			calls := cf.FuncDecl(n.(*ast.FuncDecl))
			blocks = append(blocks, calls.Blocks...)
		},
	)

	for _, block := range blocks {
		spew.Dump(block)
	}

	return nil, nil
}

func isPanicCall(n ast.Node) bool {
	expr, ok := n.(*ast.ExprStmt)
	if !ok {
		return false
	}
	call, ok := expr.X.(*ast.CallExpr)
	if !ok {
		return false
	}
	ident, ok := call.Fun.(*ast.Ident)
	if !ok {
		return false
	}
	return ident.Name == "panic"
}
