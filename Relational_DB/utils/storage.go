package storage

import (
	"fmt"
)

type Table struct {
	Rows int
	Cols int
}

// GetRows returns the number of rows
func (a *Table) GetRows() int {
	return a.Rows
}

// GetCols returns the number of columns
func (a *Table) GetCols() int {
	return a.Cols
}
