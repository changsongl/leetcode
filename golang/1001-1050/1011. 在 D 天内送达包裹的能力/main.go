package main

import (
	"fmt"
)

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}

func shipWithinDays(weights []int, D int) int {
	// 确定二分查找边界
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}

	// for 二分查找
	for left < right {
		day, pw := 1, 0
		mid := left + (right-left)/2
		for _, w := range weights {
			if pw+w > mid {
				pw = 0
				day++
			}
			pw += w
		}

		if day <= D {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
