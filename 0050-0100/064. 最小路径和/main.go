package main

// classic dp
func minPathSum(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if row == 0 && col == 0 {
				continue
			} else if row == 0 {
				grid[0][col] += grid[0][col-1]
			} else if col == 0 {
				grid[row][0] += grid[row-1][0]
			} else {
				l, t := grid[row][col-1], grid[row-1][col]
				if l > t {
					grid[row][col] += t
				} else {
					grid[row][col] += l
				}
			}
		}
	}

	return grid[rows-1][cols-1]
}
