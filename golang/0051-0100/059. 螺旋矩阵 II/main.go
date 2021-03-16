package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	var (
		left, right, top, bot = 0, n - 1, 0, n - 1
		i                     = 1
	)

	for left <= right && top <= bot {
		for col := left; col <= right; col++ {
			matrix[top][col] = i
			i++
		}

		for row := top + 1; row <= bot; row++ {
			matrix[row][right] = i
			i++
		}

		if left < right && top < bot {
			for col := right - 1; col >= left; col-- {
				matrix[bot][col] = i
				i++
			}

			for row := bot - 1; row > top; row-- {
				matrix[row][left] = i
				i++
			}
		}

		left++
		right--
		top++
		bot--
	}

	return matrix
}
