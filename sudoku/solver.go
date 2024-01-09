package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	board := os.Args[1:]
	if !verifyboard(board) {
		for _, ch := range "Error\n" {
			z01.PrintRune(ch)
		}
		return
	}
	sol := solve(board, 0, []string{})
	if len(sol) != 9 {
		for _, ch := range "Error\n" {
			z01.PrintRune(ch)
		}
		return
	}
	printboard(sol)
}

func issafe(board []string, index int, num int) bool {
	ch := byte(num + 48)
	row := index / 9
	column := index % 9
	for i := 0; i < 9; i++ {
		if board[i][column] == ch && i != row {
			return false
		}
		if board[row][i] == ch && i != column {
			return false
		}
	}
	rowQuadStart := row - row%3          // Sets value to either 0, 3 or 6 (left, middle or right)
	columnQuadStart := column - column%3 // Sets value to either 0, 3 or 6 (top, middle or bottom)
	for i := rowQuadStart; i < rowQuadStart+3; i++ {
		for j := columnQuadStart; j < columnQuadStart+3; j++ {
			if board[i][j] == ch && (row != i || column != j) {
				return false
			}
		}
	}
	return true
}

func verifyboard(board []string) bool {
	if len(board) != 9 {
		return false
	}
	hints := 0 // Hints (numbers given in the problem) counter
	for _, str := range board {
		if len(str) != 9 {
			return false
		}
	}
	for i, str := range board {
		for j, ch := range str {
			if ch != '.' && (ch < '1' && ch > '9') {
				return false
			}
			if ch >= '1' && ch <= '9' {
				if !issafe(board, i*9+j, int(ch-48)) {
					return false
				}
				hints++
			}
		}
	}
	if hints < 17 {
		return false
	}
	return true
}

func nextindex(board []string, index int) int {
	for index < 81 {
		ch := board[index/9][index%9]
		if ch == '.' {
			return index
		}
		index++
	}
	return index
}

func solve(board []string, index int, sol []string) []string {
	index = nextindex(board, index) // Unified index (0-80) for both rows and columns: index / 9 = row index (0-8), index % 9 = column index (0-8)
	if index > 80 {                 // If index is out of bounds, then we have obtained a solution
		if len(sol) == 0 { // First solution found, copy the board and memorize it
			sol = make([]string, 9)
			copy(sol, board)
			return sol
		}
		// Second solution found, invalid sudoku problem, exit immediately
		return []string{"F"}
	}
	for num := 1; num < 10; num++ { //if !issafe, increment
		// Only insert the value if it's safe
		if issafe(board, index, num) {
			board[index/9] = board[index/9][:index%9] + string(byte(num+48)) + board[index/9][index%9+1:]
			newsol := solve(board, index, sol)
			if len(newsol) == 9 {
				return newsol
			}
			board[index/9] = board[index/9][:index%9] + "." + board[index/9][index%9+1:]
		}
	}
	return sol //after its fully solved, index 80 = '.'; flow of control goes back to newsol and the board reverts to its original state
}

func printboard(board []string) {
	for _, str := range board {
		for i, ch := range str {
			z01.PrintRune(ch)
			if i < 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
