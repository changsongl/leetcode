package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}

func majorityElementHash(nums []int) int {
	dict := make(map[int]int)
	threshold := int(math.Ceil(float64(len(nums)) / 2))
	for i := 0; i < len(nums); i++ {
		num, _ := dict[nums[i]]
		dict[nums[i]] = num + 1
		if num+1 == threshold {
			return nums[i]
		}
	}

	return 0
}

func majorityElement(nums []int) int {
	vote, x := 0, 0
	for i := 0; i < len(nums); i++ {
		if vote == 0 {
			x = nums[i]
		}
		if nums[i] == x {
			vote += 1
		} else {
			vote += -1
		}
	}
	return x
}
