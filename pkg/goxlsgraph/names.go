package goxlsgraph

import (
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// DefinedNames maps excel defined names to cell addresses
type DefinedNames map[string]string

// ConstructNamesMap creates defined names map
func (d *DefinedNames) ConstructNamesMap(f *excelize.File) {
	for _, name := range f.GetDefinedName() {
		t := strings.Replace(name.RefersTo, "$", "", -1)
		axis := t[strings.IndexByte(t, '!')+1:]
		(*d)[name.Name] = axis
	}
}
