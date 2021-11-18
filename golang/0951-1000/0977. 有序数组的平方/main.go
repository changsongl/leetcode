package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-3, 2, 4}))
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	l, r, i := 0, n-1, n-1

	ret := make([]int, n)

	for l <= r {
		lNum, rNum := abs(nums[l]), abs(nums[r])
		if lNum > rNum {
			ret[i] = lNum * lNum
			l++
		} else {
			ret[i] = rNum * rNum
			r--
		}
		i--
	}

	return ret
}

func abs(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}
