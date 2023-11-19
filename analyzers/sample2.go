package analyzers

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var SampleAnalyzer = &analysis.Analyzer{
	Name: "sample",
	Doc:  "sample analyzer",
	Run:  run2,
	Requires: []*analysis.Analyzer{
		FindTestFunc,
	},
}

func run2(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok {
				if obj, ok := pass.TypesInfo.Defs[decl.Name].(*types.Func); ok {
					isTestFunc := pass.ImportObjectFact(obj, new(TestFact))
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
