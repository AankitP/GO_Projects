package main

import (
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Board Functions and struct

// Board Struct
type Board struct {
	board [][]int
}

// Function to make a new board
func newBoard() *Board {
	board := make([][]int, 6)
	for i := 0; i < 6; i++ {
		board[i] = make([]int, 7)
		for b := 0; b < 7; b++ {
			board[i][b] = 0
		}
	}
	return &Board{
		board: board,
	}
}

// Function to find the first row in a column that is empty
// must traverse the board backwards
func (t *Board) getEmptyRow(col int) (int, error) {

	if col >= 7 || col < 0 {
		return -1, fmt.Errorf("index out of range")
	}

	for i := 5; i >= 0; i-- {
		if t.board[i][col] == 0 {
			return i, nil
		}
	}

	return -1, fmt.Errorf("Column is full, cannot add more values")
}

// function to add values to board
func (t *Board) addValToCol(val, col int) error {

	fmt.Println(val, " ", col)

	if val != 1 && val != 2 {
		fmt.Println("Values must be either 1 or 2")
		err := fmt.Errorf("Incorrect Value")
		return err
	}

	emptRow, err := t.getEmptyRow(col)
	fmt.Println(emptRow)

	if err == nil {
		t.board[emptRow][col] = val
		return nil
	}
	fmt.Println(err)
	return err
}

// Function to print the whole board by row
func (t *Board) printBoard() {
	for i := 0; i < 6; i++ {
		fmt.Println(t.board[i])
	}
	fmt.Println()
}

// Get user input
func getUserInput() int {
	var a int
	for {
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer:")
			// Clear the input buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		break
	}
	return a
}
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	fmt.Println("Welcome to connect 4")
	a := newBoard()
	a.printBoard()

	// main loop to run program
	running := true
	for running {
		fmt.Println("enter 1 to add value \nenter 0 to exit")
		userInput := getUserInput()
		fmt.Printf("%T : %d\n", userInput, userInput)
		switch userInput {
		case 0:
			fmt.Println("Exiting")
			running = false

		case 1:
			fmt.Println("type value")
			value := getUserInput()
			fmt.Println("type column")
			column := getUserInput()
			fmt.Println("entered values : ", value, column)
			fmt.Println("Adding Value to Column")
			a.addValToCol(value, column)

		}
		a.printBoard()
	}
}
