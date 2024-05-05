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

var fileset *token.FileSet = token.NewFileSet()

func ParseFile() {
	fmt.Println("Hello, World!")
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
		fmt.Println("position of fileset: %v", fileset.Position(n.Pos()))
		switch expr := n.(type) {
		case *ast.BinaryExpr:
			typ := typesinfo.TypeOf(expr)
			fmt.Println("ast.BinaryExpr: %v", typ)
		case *ast.Ident:
			typ := typesinfo.TypeOf(expr)
			fmt.Println("ast.Ident: %v", typ)
		case *ast.BasicLit:
			typ := typesinfo.TypeOf(expr)
			fmt.Println("ast.BasicLit: %v", typ)
		case *ast.CallExpr:
			typ := typesinfo.TypeOf(expr)
			fmt.Println("ast.CallExpr: %v", typ)
		default:
			fmt.Printf("other Node type: %T\n", n)
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
	var typ = types.Typ
	print(typ)
	ParseFile()
	//ParseExpr()
}
