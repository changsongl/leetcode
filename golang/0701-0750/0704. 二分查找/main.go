package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (r-l)/2 + l
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}

	return -1
}
