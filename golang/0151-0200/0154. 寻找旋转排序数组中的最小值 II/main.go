package main

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	if nums[low] < nums[high] {
		return nums[low]
	}

	mid := (low + high) / 2
	for low != mid {
		if nums[mid] == nums[low] && nums[mid] == nums[high] {
			return min(findMin(nums[low:mid]), findMin(nums[mid:high]))
		}

		if nums[mid] >= nums[low] {
			low = mid
		} else if nums[high] >= nums[mid] {
			high = mid
		}

		mid = (high + low) / 2
	}

	return nums[high]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
