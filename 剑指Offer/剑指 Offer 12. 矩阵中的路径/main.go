package main

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if checkExist(word, row, col, board) {
				return true
			}
		}
	}
	return false
}

func checkExist(word string, row, col int, board [][]byte) bool {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) {
		return false
	} else if word[0] != board[row][col] {
		return false
	} else if len(word) == 1 {
		return true
	}

	letter := board[row][col]
	board[row][col] = ' '
	if checkExist(word[1:], row+1, col, board) ||
		checkExist(word[1:], row-1, col, board) ||
		checkExist(word[1:], row, col+1, board) ||
		checkExist(word[1:], row, col-1, board) {
		return true
	}

	board[row][col] = letter
	return false
}
