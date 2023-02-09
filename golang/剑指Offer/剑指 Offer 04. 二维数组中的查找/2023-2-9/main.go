package main

func findNumberIn2DArray1(matrix [][]int, target int) bool {
	var findNum func(x, y int) bool
	findNum = func(x, y int) bool {
		if x >= len(matrix) || y >= len(matrix[x]) || matrix[x][y] > target {
			return false
		} else if matrix[x][y] == target {
			return true
		} else if y == 0 && findNum(x+1, y) {
			return true
		} else if findNum(x, y+1) {
			return true
		}

		return false
	}

	return findNum(0, 0)
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	row, col := 0, cols-1

	for row >= 0 && row < rows && col >= 0 && col < cols {
		num := matrix[row][col]

		if num == target {
			return true
		} else if num > target {
			col--
		} else {
			row++
		}

	}

	return false
}
