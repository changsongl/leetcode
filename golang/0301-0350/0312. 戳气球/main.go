package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}), 167)
}

func maxCoins(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	for i := 0; i < n; i++ {
		val[i+1] = nums[i]
	}
	val[0], val[n+1] = 1, 1

	rec := make([][]int, n+2)
	for i := 0; i < len(rec); i++ {
		rec[i] = make([]int, n+2)
		for j := 0; j < n+2; j++ {
			rec[i][j] = -1
		}
	}

	return solve(0, n+1, val, rec)
}

func solve(left, right int, val []int, rec [][]int) int {
	if left >= right-1 {
		return 0
	}

	if rec[left][right] != -1 {
		return rec[left][right]
	}

	ans := 0
	for i := left + 1; i < right; i++ {
		ans = max(ans, val[left]*val[i]*val[right]+solve(left, i, val, rec)+solve(i, right, val, rec))
	}

	rec[left][right] = ans
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
