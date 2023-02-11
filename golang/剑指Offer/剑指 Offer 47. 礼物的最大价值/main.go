package main

func maxValue(grid [][]int) int {
	for i := range grid {
		for j := range grid[i] {
			var top, left int
			if i > 0 {
				top = grid[i-1][j]
			}

			if j > 0 {
				left = grid[i][j-1]
			}

			grid[i][j] = max(grid[i][j]+top, grid[i][j]+left)
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
