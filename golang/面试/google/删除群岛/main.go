package main

import "fmt"

func main() {
	m := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 1},
	}

	RemoveIslandDFS(m)

	fmt.Println(m)
}

func RemoveIslandDFS(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}

	cols := len(matrix[0])
	if cols == 0 {
		return
	}

	var check func(row, col int)

	check = func(row, col int) {
		if row < 0 || col < 0 || row >= rows || col >= cols {
			return
		} else if matrix[row][col] != 1 {
			return
		}

		matrix[row][col] = 2
		check(row+1, col)
		check(row-1, col)
		check(row, col+1)
		check(row, col-1)
	}

	for i := 0; i < rows; i++ {
		check(i, 0)
		check(i, cols-1)
	}

	for j := 1; j < cols-1; j++ {
		check(0, j)
		check(rows-1, j)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] != 0 {
				matrix[row][col] -= 1
			}
		}
	}
}
