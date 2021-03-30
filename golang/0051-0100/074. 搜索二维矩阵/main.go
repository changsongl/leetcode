package main

import "fmt"

func main() {
	//fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {30, 31, 32, 33}}, 13))
	fmt.Println(searchMatrix([][]int{{1, 3}}, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	row := 0

	for row < m {
		if matrix[row][n-1] >= target {
			break
		}
		row++
	}

	if row == m {
		return false
	}

	s := matrix[row]
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if s[m] == target {
			return true
		} else if s[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}
