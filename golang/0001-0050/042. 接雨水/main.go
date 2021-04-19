package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	left, right := make([]int, n), make([]int, n)
	for i := 1; i < n-1; i++ {
		left[i] = max(left[i-1], height[i-1])
		right[n-i-1] = max(right[n-i], height[n-i])
	}

	area := 0
	for i := 1; i < n-1; i++ {
		if height[i] < left[i] && height[i] < right[i] {
			area += min(left[i], right[i]) - height[i]
		}
	}

	return area
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
