package main

import (
	"flag"
	"fmt"
	"golang.org/x/tools/go/analysis"
)

func main() {
	fmt.Println("Hello, World!")
}

var a = &analysis.Analyzer{
	Name:             "awesome analyzer",
	Doc:              "",
	URL:              "https://github.com/0num4",
	Flags:            flag.FlagSet{},
	Run:              Run,
	RunDespiteErrors: false,
	Requires:         nil,
	ResultType:       nil,
	FactTypes:        nil,
}

//func New(conf any) []*analysis.Analyzer {
//	analyzer := &analysis.Analyzer{}
//	// return analyzer
//}
//
//func Test(t *testing.T) {
//	testdata := analysistest.TestData()
//	analysistest.Run(t, testdata)
//}

func Run(pass *analysis.Pass) (any, error) {
	// Run the analyzer
	fmt.Println("Running the analyzer")
	return nil, nil
}
