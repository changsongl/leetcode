package main

import "fmt"

func main() {
	fmt.Println(findNumberIn2DArray(
		[][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}, 5,
	))
}

// 执行用时：24 ms, 在所有 Go 提交中击败了96.60%的用户
// 内存消耗：6.6 MB, 在所有 Go 提交中击败了71.37%的用户
//
// 思路是从左上角出发, 如果从右上角开始，if条件可少一个
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	r, c := 0, 0
	for r >= 0 && r < rows && c >= 0 && c < cols {
		num := matrix[r][c]
		if num == target {
			return true
		}

		if num > target {
			c--
		} else if c+1 < cols && matrix[r][c+1] <= target {
			c++
		} else {
			r++
		}

	}

	return false
}
