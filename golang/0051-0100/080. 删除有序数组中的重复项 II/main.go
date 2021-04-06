package main

import "fmt"

func main() {
	a := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(a), a)
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	slow, fast := 2, 2
	for fast < n {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}

	return slow
}

func removeDuplicatesMine(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		checkNum := nums[i]
		repeatTime := 1
		for i+1 < len(nums) && checkNum == nums[i+1] {
			repeatTime++
			i++
		}

		for j := 0; j < repeatTime && j < 2; j++ {
			nums[index] = checkNum
			index++
		}
	}

	return index
}
