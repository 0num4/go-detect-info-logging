package main

import (
	"flag"
	"fmt"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func main() {
	fmt.Println("Hello, World!")
}

var a = &analysis.Analyzer{
	Name:             "awesomeAnalyzer",
	Doc:              "",
	URL:              "https://github.com/0num4",
	Flags:            flag.FlagSet{},
	Run:              Run,
	RunDespiteErrors: false,
	Requires:         nil,
	ResultType:       nil,
	FactTypes:        nil,
}

//	func New(conf any) []*analysis.Analyzer {
//		analyzer := &analysis.Analyzer{}
//		// return analyzer
//	}
func AnalyzerTest(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, a, "a")
}

func Run(pass *analysis.Pass) (any, error) {
	// Run the analyzer
	fmt.Println("Running the analyzer")
	return nil, nil
}
