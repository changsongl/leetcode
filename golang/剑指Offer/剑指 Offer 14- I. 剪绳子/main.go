package main

import (
	"fmt"
	"math"
)

func main() {
	cuttingRope(9)
	for i := 2; i < 10; i++ {
		fmt.Println(i, cuttingRope(i), cuttingRopeMath(i))
	}
}

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	dp := make([]int, n+1)
	dp[2] = 2

	for i := 3; i < n+1; i++ {
		for j := 2; j < i; j++ {
			dp[i] = int(math.Max(math.Max(float64(dp[j]*(i-j)), float64(j*(i-j))), float64(dp[i])))
		}
	}

	return dp[n]
}

// 数学解法当 n1+n2+...nn = n, Max(n1*n2*...nn)，分段最接近3
func cuttingRopeMath(n int) int {
	if n <= 3 {
		return n - 1
	}

	a, b := n/3, n%3
	if b == 0 {
		return int(math.Pow(3.0, float64(a)))
	} else if b == 1 {
		return int(math.Pow(3.0, float64(a-1)) * 4)
	} else {
		return int(math.Pow(3.0, float64(a)) * 2)
	}
}
