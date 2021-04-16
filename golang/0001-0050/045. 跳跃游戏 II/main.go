package main

import (
	"fmt"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)

	for i := 0; i < n-1; i++ {
		nextStep := dp[i] + 1
		for step := 1; step <= nums[i] && i+step < n; step++ {
			if dp[i+step] == 0 || nextStep < dp[i+step] {
				dp[i+step] = nextStep
			}
		}
	}

	return dp[n-1]
}
