package main

import (
	"github.com/yot0201412/go-analyzer/analyzers"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		analyzers.FindTestFunc,
		analyzers.SampleAnalyzer,
	)
}
