package main

func findRepeatNumberHash(nums []int) int {
	h := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if h[nums[i]] {
			return nums[i]
		}

		h[nums[i]] = true
	}
	return -1
}

// 鸽巢原理
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for num != i {
			if nums[num] == num {
				return num
			}

			nums[num], nums[i] = nums[i], nums[num]
		}
	}
	return -1
}

// 负数下标法
func findRepeatNumberNeg(nums []int) int {
	for i := 0; i < len(nums); i++ {
		nums[i] += 1
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num *= -1
		}
		num -= 1

		if nums[num] < 0 {
			return num
		}

		nums[num] *= -1
	}
	return -1
}
