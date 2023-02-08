package main

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
