package main

import "fmt"

func main() {
	s := []int{0, 1, 0, 3, 12}
	moveZeroes(s)
	fmt.Println(s)
}

//func moveZeroes(nums []int) {
//	n := len(nums)
//	left, right := 0, n
//	for i := 0; i < n; i++ {
//		if nums[i] == 0 {
//			right--
//		} else {
//			nums[left] = nums[i]
//			left++
//		}
//	}
//
//	for right < n {
//		nums[right] = 0
//		right++
//	}
//}

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
