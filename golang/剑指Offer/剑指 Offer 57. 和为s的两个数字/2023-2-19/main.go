package main

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			return []int{nums[l], nums[r]}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}

	return nil
}
