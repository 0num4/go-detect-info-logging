package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

var typesinfo = types.Info{
	Types:        nil,
	Instances:    nil,
	Defs:         nil,
	Uses:         nil,
	Implicits:    nil,
	Selections:   nil,
	Scopes:       nil,
	InitOrder:    nil,
	FileVersions: nil,
}

func ParseFile() {
	fmt.Println("Hello, World!")
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, "testdata/src/sampleSrc/sampleSrc.go", nil, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(fileset)
	fmt.Println(file)

	cfg := &types.Config{
		Context:                  nil,
		GoVersion:                "1.22",
		IgnoreFuncBodies:         false,
		FakeImportC:              false,
		Error:                    nil,
		Importer:                 importer.Default(),
		Sizes:                    nil,
		DisableUnusedImportCheck: false,
	}
	p, err := cfg.Check("main", fileset, []*ast.File{file}, &typesinfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("package is")
	fmt.Println(p)
	// ast.Node == ast.Fileなのかな
	ast.Inspect(file, inspectInner)
}

func inspectInner(n ast.Node) bool {

	if n != nil {
		fmt.Println("inspect inner")
		fmt.Println("position: %v", n.Pos())
		//fmt.Println("type: %v", typesinfo.TypeOf(n.(type)))
		switch expr := n.(type) {
		case *ast.BinaryExpr:

			println("types: %v", typesinfo.TypeOf(expr))
		}
	}
	return true
}

//func InspectSrc(file ) {
//	//astnode := ast.Node()
//
//}

func ParseExpr() {
	fmt.Println("parse_expr")
	expr, err := parser.ParseExpr("1*2+3")
	if err != nil {
		panic(err)
	}
	fmt.Println(expr)

}

func main() {
	ParseFile()
	//ParseExpr()
}
