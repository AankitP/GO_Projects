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

	if val != 1 && val != 2 {
		fmt.Println("Values must be either 1 or 2")
		err := fmt.Errorf("Incorrect Value")
		return err
	}

	emptRow, err := t.getEmptyRow(col)

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
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		break
	}
	return a
}

func (t *Board) connect4HorizontalChecker() (bool, int) {
	fmt.Println("Checking Board Horizontally")
	for i := 0; i < 6; i++ {
		a := t.board[i]
		fmt.Println(a)

		checkVal := a[0]
		totalCheckVal := 0
		for j := 0; j<7; j++{
			if a[j] == checkVal{
				totalCheckVal++
			}else{
				totalCheckVal = 1
				checkVal = a[j]
			}
			if(totalCheckVal == 4 && checkVal != 0){
				return true, checkVal
			}
		}
	}
	return false, -1
}

func (t *Board) connect4VerticalChecker() (bool, int) {
	fmt.Println("Checking Board Vertically")

	return false, -1
}

func (t *Board) connect4DiagonalChecker() (bool, int) {
	fmt.Println("Checking Board Diagonally")

	return false, -1
}

func (t *Board) connect4Checker() (bool, int) {
	fmt.Println("Checking Board")

	Horizontal, Winner := t.connect4HorizontalChecker()
	if Horizontal{
		return true, Winner
	}
	Vertical, Winner := t.connect4VerticalChecker()
	if Vertical{
		return true, Winner
	}
	Diagonal, Winner := t.connect4DiagonalChecker()
	if Diagonal {
		return true, Winner
	}

	return false, -1
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	fmt.Println("Welcome to connect 4")
	a := newBoard()
	a.printBoard()

	// main loop to run program
	running := true
	for running {
		fmt.Println("enter 0 to exit\nenter 1 to add value\nenter 2 to reset game")
		userInput := getUserInput()
		switch userInput {
		case 0:
			fmt.Println("Exiting")
			running = false

		case 1:
			fmt.Println("type value")
			value := getUserInput()
			fmt.Println("type column")
			column := getUserInput()
			// fmt.Println("Adding Value to Column")
			a.addValToCol(value, column)

		case 2:
			fmt.Println("Clearing Board")
			a = newBoard()
			fmt.Println("Board cleared")

		case 3:
			fmt.Println("Checking Board")
			fmt.Println(a.connect4HorizontalChecker())
		}
		a.printBoard()
	}
}
