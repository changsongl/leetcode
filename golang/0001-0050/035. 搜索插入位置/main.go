package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3}, 2))
}

func searchInsert(nums []int, target int) int {
	low, high, checkNum, mid := 0, len(nums)-1, 0, 0
	for low <= high {
		mid = (low + high) / 2
		checkNum = nums[mid]

		if checkNum == target {
			return mid
		} else if checkNum > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	// 点睛之笔，省略最后的if判断
	return low
}
