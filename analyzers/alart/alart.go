package alart

import (
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
	result := pass.ResultOf[findtestfunc.Analyzer].(*findtestfunc.FindTestFuncResult)
	if len(result.HasFactFiles) == 0 {
		return nil, nil
	}
	for _, f := range pass.Files {
		// アナライザーの結果はパッケージ別に返ってくるっぽいが、調査
		_, ok := result.HasFactFiles[f]
		if !ok {
			continue
		}
		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok {
				if obj, ok := pass.TypesInfo.Defs[decl.Name].(*types.Func); ok {
					isTestFunc := pass.ImportObjectFact(obj, new(findtestfunc.FindTestFuncFact))
					if isTestFunc {
						pass.Reportf(decl.Pos(), "This is test func.")
					}

				}
			}
		}

	}
	return nil, nil
}
