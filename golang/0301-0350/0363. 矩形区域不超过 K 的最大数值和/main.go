package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
	fmt.Println(maxSumSubmatrix([][]int{{2, 2, -1}}, 3))
	fmt.Println(maxSumSubmatrix([][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}, 8))
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 {
		return 0
	}

	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			top, left, topLeft := 0, 0, 0
			if i > 0 {
				top = dp[i-1][j]
			}
			if j > 0 {
				left = dp[i][j-1]
			}
			if i > 0 && j > 0 {
				topLeft = dp[i-1][j-1]
			}
			dp[i][j] = top + left - topLeft + matrix[i][j]
		}
	}

	max := math.MinInt64
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for i2 := i; i2 < n; i2++ {
				for j2 := j; j2 < m; j2++ {
					topX, topY, botX, botY := i, j, i2, j2
					if botX < n && botY < m {
						value := matrix[i][j]
						if i != i2 || j != j2 {
							if topX == 0 && topY == 0 {
								value = dp[botX][botY]
							} else if topX == 0 {
								value = dp[botX][botY] - dp[botX][topY-1]
							} else if topY == 0 {
								value = dp[botX][botY] - dp[topX-1][botY]
							} else {
								value = dp[botX][botY] - dp[topX-1][botY] - dp[botX][topY-1] + dp[topX-1][topY-1]
							}
						}

						if value > max && value <= k {
							max = value
						}
					}
				}
			}
		}
	}
	return max
}
