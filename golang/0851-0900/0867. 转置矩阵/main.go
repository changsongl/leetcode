package main

func transpose(matrix [][]int) [][]int {
	cols, rows := len(matrix[0]), len(matrix)
	newMat := make([][]int, 0, cols)
	for col := 0; col < cols; col++ {
		newRow := make([]int, rows)
		for row := 0; row < rows; row++ {
			newRow[row] = matrix[row][col]
		}
		newMat = append(newMat, newRow)
	}

	return newMat
}
