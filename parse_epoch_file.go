package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func ParseFile() {
	fmt.Println("Hello, World!")
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, "testdata/src/sampleSrc/sampleSrc.go", nil, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(fileset)
	fmt.Println(file)
}

func main() {
	ParseFile()
}
