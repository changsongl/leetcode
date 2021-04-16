package main

import "fmt"

func main() {
	fmt.Println(search([]int{7, 8, 9, 1, 2, 3, 4, 5, 6}, 1))
	fmt.Println(search([]int{7, 8, 9, 1, 2, 3, 4, 5, 6}, 2))
	fmt.Println(search([]int{7, 8, 9, 1, 2, 3, 4, 5, 6}, 5))
	fmt.Println(search([]int{7, 8, 9, 1, 2, 3, 4, 5, 6}, 10))
}

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1

	for l <= r {
		mid := l + (r-l)/2

		// 左边为正序
		if nums[mid] == target {
			return mid
		} else if nums[0] <= nums[mid] { // 左边为正序
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // 右边正序
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
