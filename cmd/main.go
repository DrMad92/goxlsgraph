package main

import (
	"fmt"

	"github.com/DrMad92/goxlsgraph/pkg/goxlsgraph"
)

func main() {

	graph, err := goxlsgraph.ExcelToGraph("test.xlsx", "Sheet1")
	if err != nil {
		panic(err)
	}
	// var definedNames = make(goxlsgraph.DefinedNames)
	// definedNames.ConstructNamesMap(f)

	for k, v := range *graph {
		fmt.Print(k)
		fmt.Println(v)
	}
}
