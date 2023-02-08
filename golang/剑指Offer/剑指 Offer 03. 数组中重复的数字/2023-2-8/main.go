package main

func findRepeatNumber1(nums []int) int {
	m := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, exist := m[num]; exist {
			return num
		}

		m[num] = struct{}{}
	}

	return -1
}

func findRepeatNumber2(nums []int) int {
	m := make([]int, 100000)

	for _, num := range nums {
		if m[num] > 0 {
			return num
		}

		m[num]++
	}

	return -1
}

func findRepeatNumber3(nums []int) int {
	var i int

	for i < len(nums) {
		if i == nums[i] {
			i++
			continue
		}

		if nums[i] == nums[nums[i]] {
			return nums[i]
		}

		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}

	return -1
}

func main() {
	findRepeatNumber3([]int{2, 3, 1, 0, 2, 5, 3})
}
