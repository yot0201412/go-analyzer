package findtestfunc

import (
	"go/ast"
	"go/types"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "FindTestFunc",
	Doc:  "Find test func",
	Run:  run,
	FactTypes: []analysis.Fact{
		new(FindTestFuncFact),
	},
	Requires:   []*analysis.Analyzer{},
	ResultType: reflect.TypeOf(new(FindTestFuncResult)),
}

func run(pass *analysis.Pass) (interface{}, error) {
	res := &FindTestFuncResult{}
	hasFactFiles := make(map[*ast.File]struct{})
	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok && strings.HasPrefix(decl.Name.Name, "Test") {
				if obj, ok := pass.TypesInfo.Defs[decl.Name].(*types.Func); ok {
					hasFactFiles[f] = struct{}{}
					pass.ExportObjectFact(obj, new(FindTestFuncFact))
				}
			}
		}
	}
	res.HasFactFiles = hasFactFiles
	return res, nil
}

type FindTestFuncFact struct {
	// IsTest bool
}

func (FindTestFuncFact) AFact() {}

func (*FindTestFuncFact) String() string { return "findTestFuncFact" }

type FindTestFuncResult struct {
	HasFactFiles map[*ast.File]struct{}
}
