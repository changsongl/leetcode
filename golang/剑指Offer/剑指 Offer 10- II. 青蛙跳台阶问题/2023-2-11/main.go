package main

import "fmt"

func numWays(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a
}

func numWays1(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}

	return dp[n]
}

func main() {
	fmt.Println(numWays(7))
}
