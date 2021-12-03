package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{-4, -2, -3}, 4))
}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums) && k > 0; i++ {
		if nums[i] < 0 && i != len(nums)-1 {
			nums[i] = -nums[i]
		} else if k%2 == 1 {
			if i == 0 || nums[i] < nums[i-1] {
				nums[i] = -nums[i]
			} else {
				nums[i-1] = -nums[i-1]
			}
			k = 0
		} else {
			k = 0
		}

		k--
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
