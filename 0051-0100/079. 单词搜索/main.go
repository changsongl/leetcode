package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{{'a', 'a'}}, "aaa"))
}

func exist(board [][]byte, word string) bool {
	fl := word[0]
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == fl {
				old := board[row][col]
				board[row][col] = ' '
				if search(board, word, row, col) {
					return true
				}
				board[row][col] = old
			}
		}
	}

	return false
}

func search(board [][]byte, word string, row, col int) bool {
	if len(word) == 1 {
		return true
	}

	checkChar := word[1]

	// up
	if row > 0 && board[row-1][col] == checkChar {
		if keepSearch(board, word, row-1, col) {
			return true
		}
	}

	// down
	if row < len(board)-1 && board[row+1][col] == checkChar {
		if keepSearch(board, word, row+1, col) {
			return true
		}
	}

	// left
	if col > 0 && board[row][col-1] == checkChar {
		if keepSearch(board, word, row, col-1) {
			return true
		}
	}

	// right
	if col < len(board[row])-1 && board[row][col+1] == checkChar {
		if keepSearch(board, word, row, col+1) {
			return true
		}
	}

	return false
}

func keepSearch(board [][]byte, word string, row, col int) bool {
	old := board[row][col]
	board[row][col] = ' '
	if search(board, word[1:], row, col) {
		return true
	}
	board[row][col] = old
	return false
}
