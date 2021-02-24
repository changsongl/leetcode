package main

import "fmt"

//[4,5,6,7,0,1,2]
// 0,1,2,3,4,5,6
// 3 , 6
// 4 , 6
// 4 - (6 - 4) = 2

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	if nums[low] < nums[high] {
		return nums[low]
	}

	mid := (low + high) / 2
	for low != mid {
		if nums[mid] > nums[low] {
			low = mid
		} else if nums[high] > nums[mid] {
			high = mid
		}

		mid = (high + low) / 2
	}

	return nums[high]
}

func findMinMine(nums []int) int {
	l, r := 0, len(nums)-1
	for l != r {
		if nums[l] > nums[r] {
			if l+1 == r {
				break
			}
			l = (r + l) / 2
		} else {
			l, r = l*2-r+1, l
			if l < 0 {
				l = 0
			}
		}
	}
	return nums[r]
}
