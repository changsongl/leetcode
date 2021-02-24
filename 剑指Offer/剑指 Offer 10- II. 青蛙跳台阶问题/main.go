package main

func numWays(n int) int {
	var a, b int64 = 1, 1
	for i := 0; i < n; i++ {
		a, b = b%1000000007, (a+b)%1000000007
	}

	return int(a % 1000000007)
}
