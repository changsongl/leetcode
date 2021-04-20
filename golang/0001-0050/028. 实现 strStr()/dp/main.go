package main

import "fmt"

func main() {
	// 0 1 2 3 4 5
	// 1 3 1 5 8 1
	fmt.Println(maxCoins([]int{3, 1, 5, 8}), 167)
}

func maxCoins(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	for i := 0; i < n; i++ {
		val[i+1] = nums[i]
	}
	val[0], val[n+1] = 1, 1

	dp := make([][]int, n+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+2)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}

	return dp[0][n+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
