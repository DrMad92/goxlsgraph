package goxlsgraph

// Cell type
type Cell struct {
	IsFormula   bool
	Formula     string
	Axis        string
	Value       string
	Precendents []*Cell
	Dependents  []*Cell
}
