package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	board := make([][]rune, n)
	for i := range board {
		board[i] = []rune(strings.Repeat(".", n))
	}
	var res [][]string
	// defer printBoard(board)
	diag1 := make([]bool, 2*n)
	diag2 := make([]bool, 2*n)
	cols := make([]bool, n)
	var backtrack func(row int)

	backtrack = func(row int) {
		if row == n {
			var soln []string
			for _, r := range board {
				soln = append(soln, string(r))
			}
			res = append(res, soln)
			return
		}
		for col := 0; col < n; col++ {
			d1 := row - col + n
			d2 := row + col
			if cols[col] || diag1[d1] || diag2[d2] {
				continue
			}

			board[row][col] = 'Q'
			cols[col], diag1[d1], diag2[d2] = true, true, true
			backtrack(row + 1)
			board[row][col] = '.'
			cols[col], diag1[d1], diag2[d2] = false, false, false
		}

	}

	backtrack(0)

	return res
}
