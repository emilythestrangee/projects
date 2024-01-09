package main

import (
	"github.com/01-edu/z01"
)

var pos = [8]int{}

func issafe(queen_num int, row_pos int) bool { //comparing row pos of prev. placed queens & row pos of current queen
	for i := 0; i < queen_num; i++ { //retrieving index i(row pos) of prev queen from pos array (the one b4 queen_num)
		row_pos1 := pos[i]

		if row_pos1 == row_pos || abs(row_pos1-row_pos) == abs(queen_num-i) {
			return false
		} //queen_num is same as column pos as there is one queen w/ increase in column pos
		//i is the previous column pos as we're checking the previously placed columns
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func solve(q int) { //q: no. of queens placed
	if q == 8 { //base case- all 8 queens placed in one board
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(pos[i]) + '0') //converting zero-based index to one-based index i.e, pos[i]=0
		}
		z01.PrintRune('\n')
	} else {
		for i := 1; i < 9; i++ { //row pos of new queen
			if issafe(q, i) {
				pos[q] = i //q is kinda the same as index
				solve(q + 1)
			}
		}
	}
}

func main() {
	solve(0) //initial no. of queens placed: 0
}
