package goxlsgraph

import "github.com/360EntSecGroup-Skylar/excelize/v2"

// ExcelToGraph creates graph from passed xls filename
func ExcelToGraph(filename, sheet string) (*Graph, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	var graph = make(Graph)
	graph.generateGraphFromSheet(f, sheet)
	return &graph, nil
}
