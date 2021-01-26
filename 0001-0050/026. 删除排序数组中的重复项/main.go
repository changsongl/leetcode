package main

func removeDuplicates(nums []int) int {
	lNums := len(nums)
	if lNums < 2 {
		return lNums
	}

	index, last := 1, nums[0]

	for i := 1; i < lNums; i++ {
		if last != nums[i] {
			nums[index] = nums[i]
			last = nums[i]
			index++
		}
	}

	return index
}
