package main

func exist(board [][]byte, word string) bool {
	rows, cols, wordLen := len(board), len(board[0]), len(word)
	used := make([][]bool, rows)
	for i := range used {
		used[i] = make([]bool, cols)
	}

	var f func(wordIndex, row, col int) bool
	f = func(wordIndex, row, col int) bool {
		if wordIndex >= wordLen {
			return true
		} else if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[wordIndex] || used[row][col] {
			return false
		}

		used[row][col] = true
		defer func() {
			used[row][col] = false
		}()

		if f(wordIndex+1, row-1, col) || f(wordIndex+1, row, col-1) || f(wordIndex+1, row+1, col) || f(wordIndex+1, row, col+1) {
			return true
		}

		return false
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {

			if f(0, row, col) {
				return true
			}
		}
	}
	return false
}
