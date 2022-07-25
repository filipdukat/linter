package modules

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// NewAnalyzer returns a new *analysis.Analyzer configured for searching for obsolete libs.
func NewAnalyzer(libs []string) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "libraries",
		Doc:      "Checks for flagged libraries in a project.",
		Run:      run(libs),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(libs []string) func(pass *analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
		nodeFilter := []ast.Node{
			(*ast.ImportSpec)(nil),
		}
		inspector.Preorder(nodeFilter, func(node ast.Node) {
			importSpec := node.(*ast.ImportSpec)
			var msg = " this module is banned."

			for _, l := range libs {
				if strings.HasPrefix(strings.Trim(importSpec.Path.Value, "\""), l) {
					pass.Report(analysis.Diagnostic{
						Pos:     importSpec.Pos(),
						End:     importSpec.End(),
						Message: fmt.Sprint(importSpec.Path.Value, msg),
					})
				}
			}
		})
		return nil, nil
	}
}
