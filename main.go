package main

import (
	"flag"
	"fmt"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"testing"
)

func main() {
	fmt.Println("Hello, World!")
}

var ab = &analysis.Analyzer{
	Name:             "awesomeAnalyzer",
	Doc:              "this is awesome analyzer",
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
	analysistest.Run(t, testdata, ab, "sampleSrc")
}

func Run(pass *analysis.Pass) (any, error) {
	fmt.Println("Running the analyzer")

	// Run the analyzer

	if pass != nil {
		// awesomeAnalyzerが入ってる
		//if pass.Analyzer.Requires {
		inspect := pass.ResultOf[inspect.Analyzer]
		print(inspect)
	}
	//pass.Reportフィールド(関数を受け取る)か、pass.Reportf()を使ってエラー(Diagnostic)を報告する
	pass.Reportf(3, "This is a test error")

	fmt.Println("Running the analyzer done")
	return nil, nil
}
