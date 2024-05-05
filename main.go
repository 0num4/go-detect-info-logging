package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/ast/inspector"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var ab = &analysis.Analyzer{
	Name:             "awesomeAnalyzer",
	Doc:              "this is awesome analyzer",
	URL:              "https://github.com/0num4",
	Flags:            flag.FlagSet{},
	Run:              run,
	RunDespiteErrors: false,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	ResultType: nil,
	FactTypes:  nil,
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
		//https://zenn.dev/tenntenn/books/d168faebb1a739/viewer/9d590e#*analysis.pass%E5%9E%8B%E3%82%92%E5%85%A5%E5%8A%9B%E3%83%87%E3%83%BC%E3%82%BF%E3%81%A8%E3%81%97%E3%81%A6%E4%BD%BF%E3%81%86
		// pass.AnalyzerにはawesomeAnalyzerが入ってる
		// pass.Fsetにはtoken.Filesetが入ってる
		//　AST.fileとかもpass.Filesにある
		//　TypesInfoにはtypes.Infoとかもある。解析結果
		// if pass.Analyzer.Requires {
		//inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	}
	// pass.Reportフィールド(関数を受け取る)か、pass.Reportf()を使ってエラー(Diagnostic)を報告する
	p := token.Pos(3)
	pass.Reportf(p, "This is a test error")

	inspect := inspector.Inspector{}
	// inspect.Preorderの第一引数はフィルターするastのlist
	filter := []ast.Node{
		(*ast.Ident)(nil),
	}
	inspect.Preorder(filter, inspectInnerWrapper)

	fmt.Println(inspect)

	fmt.Println("Running the analyzer done")
	return nil, nil
}

func inspectInnerWrapper(n ast.Node) {
	inspectInner(n)
}
