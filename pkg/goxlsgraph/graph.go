package goxlsgraph

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// Graph type
type Graph map[string]Cell

// generateGraphFromSheet generates graph from excel file
func (g *Graph) generateGraphFromSheet(f *excelize.File, sheet string) error {
	rows, err := f.GetRows(sheet)
	if err != nil {
		return err
	}
	for y, row := range rows {
		for x, value := range row {
			var c = Cell{
				Value: value,
			}
			axis, err := excelize.CoordinatesToCellName(x+1, y+1)
			if err != nil {
				return err
			}
			c.Axis = axis
			formula, err := f.GetCellFormula(sheet, axis)
			if err != nil {
				return err
			}
			if formula != "" {
				c.Formula = formula
				c.IsFormula = true
			}
			(*g)[axis] = c
		}
	}
	return nil
}
