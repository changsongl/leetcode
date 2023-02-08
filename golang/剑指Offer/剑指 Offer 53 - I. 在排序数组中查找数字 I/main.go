package main

import "fmt"

func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	l, r := 0, len(nums)-1

	for l <= r {
		c := (l + r) / 2

		if nums[c] == target {
			count := 1
			for i := c - 1; i >= 0 && nums[i] == target; i-- {
				count++
			}
			for i := c + 1; i < len(nums) && nums[i] == target; i++ {
				count++
			}

			return count
		} else if nums[c] > target {
			r = c - 1
		} else {
			l = c + 1
		}
	}

	return 0
}

func search(nums []int, target int) int {
	l := binarySearch(nums, target, true)
	r := binarySearch(nums, target, false) - 1

	if l <= r && r < len(nums) && nums[l] == target && nums[r] == target {
		return r - l + 1
	}

	return 0
}

func binarySearch(nums []int, target int, lower bool) int {
	l, r, ans := 0, len(nums)-1, len(nums)
	for l <= r {
		m := (l + r) / 2
		if nums[m] > target || (lower && nums[m] >= target) {
			r = m - 1
			ans = m
		} else {
			l = m + 1
		}
	}

	return ans
}

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 8, 8, 8, 8}, 8))
}
