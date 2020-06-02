package main

//执行用时 : 4 ms , 在所有 Go 提交中击败了 96.89% 的用户
//内存消耗 : 3.6 MB , 在所有 Go 提交中击败了 57.14% 的用户
func twoSumTest(nums []int, target int) []int {
	hashes := make(map[int]int, len(nums))
	for i, num := range nums{
		i2, exists := hashes[num]
		if exists && target - num == num {
			return []int{i, i2}
		}
		hashes[num] = i
	}

	for num, i := range hashes{
		otherNum := target - num
		i2, exists := hashes[otherNum]
		if exists && i != i2{
			return []int{i, i2}
		}
	}
	return []int{}
}

//执行用时 : 60 ms , 在所有 Go 提交中击败了 9.93% 的用户
//内存消耗 : 2.9 MB , 在所有 Go 提交中击败了 100.00% 的用户
func twoSumBruteForce(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l - 1; i++{
		for j := i + 1; j < l; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
