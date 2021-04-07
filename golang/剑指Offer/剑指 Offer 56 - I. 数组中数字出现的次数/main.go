package main

import "fmt"

func main() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
}

func singleNumbers(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	bit := 1
	for (xor & bit) == 0 {
		bit = bit << 1
	}

	a, b := 0, 0
	for _, num := range nums {
		if num&bit == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
