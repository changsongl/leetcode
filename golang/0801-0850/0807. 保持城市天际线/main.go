package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	top, left := make([]int, m), make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			left[i] = max(left[i], grid[i][j])
			top[j] = max(top[j], grid[i][j])
		}
	}

	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			total += min(left[i], top[j]) - grid[i][j]
		}
	}

	return total
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
