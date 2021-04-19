package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getPermutation(3, 3))
}

func getPermutation(n int, k int) string {
	ans := ""
	count := 0

	choices := make([]int, n)
	for i := 0; i < n; i++ {
		choices[i] = i + 1
	}

	var dfs func(nums []int)
	dfs = func(nums []int) {
		subCount := getCount(len(nums) - 1)
		for i := 0; i < len(nums); i++ {
			if subCount+count >= k {
				ans += strconv.Itoa(nums[i])
				dfs(append(nums[:i], nums[i+1:]...))
				break
			} else {
				count += subCount
			}
		}
	}

	dfs(choices)
	return ans
}

func getCount(n int) int {
	if n == 0 {
		return 1
	}

	count := 1
	for i := 2; i <= n; i++ {
		count *= i
	}
	return count
}
