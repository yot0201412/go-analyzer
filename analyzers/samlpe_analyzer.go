package analyzers

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var FindTestFunc = &analysis.Analyzer{
	Name: "FindTestFunc",
	Doc:  "Find test func",
	Run:  run,
	FactTypes: []analysis.Fact{
		new(TestFact),
	},
	Requires: []*analysis.Analyzer{},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok && strings.HasPrefix(decl.Name.Name, "Test") {
				if obj, ok := pass.TypesInfo.Defs[decl.Name].(*types.Func); ok {
					// _ = obj
					pass.ExportObjectFact(obj, new(TestFact))
				}
			}
		}
	}
	return nil, nil
}

type TestFact struct {
	// IsTest bool
}

func (TestFact) AFact() {}
