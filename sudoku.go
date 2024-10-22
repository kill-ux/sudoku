package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Error() {
	z01.PrintRune('E')
	z01.PrintRune('r')
	z01.PrintRune('r')
	z01.PrintRune('o')
	z01.PrintRune('r')
	z01.PrintRune('\n')
}

func PrintString(TabSudoku []int) {
	for i, e := range TabSudoku {
		z01.PrintRune(rune(e + '0'))
		if (i+1)%9 == 0 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
}

func isValid(board []int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row*9+i] == num || board[i*9+col] == num {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i*9+j] == num {
				return false
			}
		}
	}
	return true
}

func solveSudoku(board []int) bool {
	for i := 0; i < 81; i++ {
		if board[i] == 0 {
			for num := 1; num <= 9; num++ {
				if isValid(board, i/9, i%9, num) {
					board[i] = num
					if solveSudoku(board) {
						return true
					}
					board[i] = 0
				}
			}
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 10 {
		Error()
		return
	}

	Tab := os.Args[1:]
	for _, row := range Tab {
		if len(row) != 9 {
			Error()
			return
		}
	}

	TabSudoku := make([]int, 81)
	for i, row := range Tab {
		for j, ch := range row {
			if ch >= '1' && ch <= '9' {
				TabSudoku[i*9+j] = int(ch - '0')
			} else if ch == '.' {
				TabSudoku[i*9+j] = 0
			} else {
				Error()
				return
			}
		}
	}

	if !solveSudoku(TabSudoku) {
		Error()
		return
	}

	PrintString(TabSudoku)
}
