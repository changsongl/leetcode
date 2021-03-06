package main

import (
	"fmt"
)

func main() {
	fmt.Println(nextGreaterElements([]int{7, 5, 3, 8, 9, 0}))
}

func nextGreaterElements(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return nums
	}

	stack, ans := []int{0}, make([]int, l)
	for index := 1; index < l*2; index++ {
		i := index % l
		for len(stack) != 0 {
			j := stack[len(stack)-1]
			if nums[i] > nums[j] {
				ans[j] = nums[i]
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		if index < l {
			stack = append(stack, i)
		}
	}

	for len(stack) != 0 {
		ans[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	return ans
}
