package main

//func twoSum(nums []int, target int) []int {
//	dict := make(map[int]bool, len(nums))
//	for _, num := range nums {
//		dict[num] = true
//	}
//
//	for _, num := range nums {
//		if result := dict[target-num]; result {
//			return []int{num, target - num}
//		}
//	}
//	return []int{}
//}

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i != j {
		sum := nums[i] + nums[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return []int{nums[i], nums[j]}
		}
	}

	return nil
}
