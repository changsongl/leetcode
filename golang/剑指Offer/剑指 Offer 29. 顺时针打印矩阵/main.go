package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])

	ans := make([]int, 0, m*n)
	l, r, t, b := 0, n-1, 0, m-1
	for l <= r && t <= b {
		ans = append(ans, matrix[t][l:r+1]...)

		for i := t + 1; i <= b; i++ {
			ans = append(ans, matrix[i][r])
		}

		for i := r - 1; i >= l && t != b; i-- {
			ans = append(ans, matrix[b][i])
		}

		for i := b - 1; i > t && l != r; i-- {
			ans = append(ans, matrix[i][l])
		}

		l++
		t++
		b--
		r--
	}
	return ans
}
