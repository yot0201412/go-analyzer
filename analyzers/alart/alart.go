package alart

import (
	"fmt"
	"go/ast"
	"go/types"

	"github.com/yot0201412/go-analyzer/analyzers/findtestfunc"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "alartTestFunc",
	Doc:  "Alart test func",
	Run:  run2,
	Requires: []*analysis.Analyzer{
		findtestfunc.Analyzer,
	},
}

func run2(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok {
				if obj, ok := pass.TypesInfo.Defs[decl.Name].(*types.Func); ok {
					isTestFunc := pass.ImportObjectFact(obj, new(findtestfunc.TestFact))
					fmt.Println(isTestFunc, decl.Name.Name)
					if isTestFunc {
						pass.Reportf(decl.Pos(), "This is test func.")
					}

				}
			}
		}

	}
	return nil, nil
}
