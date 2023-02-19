package main

func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l]%2 == 0 {
			nums[r], nums[l] = nums[l], nums[r]
			r--
		} else if nums[r]%2 == 1 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		} else {
			l++
			r--
		}
	}

	return nums
}
