package main

import (
	"fmt"
)

func main() {
	//fmt.Println(rob([]int{2, 3, 2}), 3)
	//fmt.Println(rob([]int{1, 2, 3, 1}), 4)
	//fmt.Println(rob([]int{1, 2, 1, 1}), 3)
	//fmt.Println(rob([]int{1, 1, 3, 6, 7, 10, 7, 1, 8, 5, 9, 1, 4, 4, 3}), 41)
	fmt.Println(rob([]int{1}), 3)
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	ans := robDp(nums[:len(nums)-1])
	if second := robDp(nums[1:]); second > ans {
		ans = second
	}
	return ans
}

func robDp(nums []int) int {
	n := len(nums)
	if n == 1 {

	}
	dp, max := make([]int, n), 0

	for i, num := range nums {
		if i < 2 {
			dp[i] = num
		} else {
			choice := num + dp[i-2]
			if i-3 >= 0 && dp[i-3] > dp[i-2] {
				choice = num + dp[i-3]
			}
			dp[i] = choice
		}

		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
