package main

import "fmt"

func main() {
	fmt.Println(search([]int{2, 5, 7, 0, 0, 1, 2}, 1))
}

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	if left <= right {
		mid := (left + right) / 2
		isReverse := nums[left] >= nums[right]

		if nums[mid] == target {
			return true
		}

		if isReverse || target < nums[mid] {
			if found := search(nums[:mid], target); found {
				return true
			}
		}

		if isReverse || target > nums[mid] {
			if found := search(nums[mid+1:], target); found {
				return true
			}
		}
	}
	return false
}

//func search(nums []int, target int) bool {
//	if len(nums) == 0 {
//		return false
//	}
//
//	if target >= nums[0] {
//		for i := 0; i < len(nums); i++ {
//			if nums[i] == target {
//				return true
//			} else if nums[i] > target || (i != 0 && nums[i] < nums[i-1]) {
//				return false
//			}
//		}
//	} else {
//		for i := len(nums) - 1; i >= 0; i-- {
//			if nums[i] == target {
//				return true
//			} else if nums[i] < target || (i >= 1 && nums[i] < nums[i-1]) {
//				return false
//			}
//		}
//	}
//
//	return false
//}
