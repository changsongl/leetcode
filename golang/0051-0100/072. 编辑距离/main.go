package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m*n == 0 {
		return max(m, n)
	}

	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	for j := 1; j < n+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		row := make([]int, n+1)
		row[0] = i
		dp[i] = row

		for j := 1; j < n+1; j++ {
			cost := min(dp[i][j-1], dp[i-1][j])
			if word1[i-1] == word2[j-1] {
				cost = min(cost, dp[i-1][j-1]-1)
			} else {
				cost = min(cost, dp[i-1][j-1])
			}
			dp[i][j] = cost + 1
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
