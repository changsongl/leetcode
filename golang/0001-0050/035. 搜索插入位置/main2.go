package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3}, 2), 1)

	fmt.Println(searchInsert([]int{1, 3}, 4), 2)

	fmt.Println(searchInsert([]int{1, 3}, 0), 0)
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}

	return l
}
