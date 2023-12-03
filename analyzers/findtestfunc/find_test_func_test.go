package findtestfunc_test

import (
	"fmt"
	"testing"

	"github.com/yot0201412/go-analyzer/analyzers/findtestfunc"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	fmt.Println(testdata)
	analysistest.Run(t, testdata, findtestfunc.Analyzer, "sample")
}
