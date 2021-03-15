package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1}}))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var (
		rows, cols            = len(matrix), len(matrix[0])
		ans                   = make([]int, 0, rows*cols)
		left, right, top, bot = 0, cols - 1, 0, rows - 1
	)

	for left <= right && top <= bot {
		ans = append(ans, matrix[top][left:right+1]...)

		for row := top + 1; row <= bot; row++ {
			ans = append(ans, matrix[row][right])
		}

		if left < right && top < bot {
			for col := right - 1; col >= left; col-- {
				ans = append(ans, matrix[bot][col])
			}

			for row := bot - 1; row > top; row-- {
				ans = append(ans, matrix[row][left])
			}
		}

		left++
		right--
		top++
		bot--
	}

	return ans
}
