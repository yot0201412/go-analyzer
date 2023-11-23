package main

import (
	"github.com/yot0201412/go-analyzer/analyzers/alart"
	"github.com/yot0201412/go-analyzer/analyzers/findtestfunc"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		findtestfunc.Analyzer,
		alart.Analyzer,
	)
}
