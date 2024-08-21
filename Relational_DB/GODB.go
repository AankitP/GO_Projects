package main

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////////////

// Table Struct
type Table struct {
	Rows int
	Cols int
	Tab  [][]string
}

// Function to create a new table, and populate it with "NotPopulated"
	// Reason to populate it is for testing purposes, will be removed later
func newTable(rows, cols int) *Table {
	tab := make([][]string, rows)
	for i := 0; i < rows; i++ {
		tab[i] = make([]string, cols)
		for b := 0; b < cols; b++{
			tab[i][b] = "NotPopulated"
		}
	}
	return &Table{
		Rows: rows,
		Cols: cols,
		Tab:  tab,
	}
}

// GetRows returns the number of rows
func (a *Table) GetRows() int {
	return a.Rows
}

// GetCols returns the number of columns
func (a *Table) GetCols() int {
	return a.Cols
}

// Function to set specific values
func (t *Table) setVal(row, col int, value string) error {
	if row < 0 || row > t.Rows || col < 0 || col > t.Cols {
		return fmt.Errorf("Index out of range")
	}
	t.Tab[row][col] = value
	return nil
}

// function to get specific Values
func (t *Table) getVal(row, col int) (string, error) {
	if row < 0 || row >= t.Rows || col < 0 || col >= t.Cols {
		return "", fmt.Errorf("index out of range")
	}
	return t.Tab[row][col], nil
}

// function to get Specific Row
func (t *Table) getRow(row int) ([]string, error){
	if row < 0 || row >= t.Rows{
		return nil, fmt.Errorf("Row is out of range")
	}
	return t.Tab[row], nil
}

// Function to print the whole table at once
func (t *Table) printTable(){
	for i := 0; i < t.Rows; i++{
		fmt.Println(t.Tab[i])
	}
}

////////////////////////////////////////////////////////////////////////////////////////

func main() {
	a := newTable(3, 4)      // Calling method to create new table
	fmt.Println(a.GetRows()) // Calling the method to get the number of rows
	fmt.Println(a.GetCols()) // Calling the method to get the number of columns

	a.setVal(2, 2, "I am here")
	// b,_ := a.getVal(2,2)
	// a.printTable()
	fmt.Println(a.getRow(2))
}
