package main

import "fmt"

func main() {
	fmt.Println(countBits(5))
}

func countBits(num int) []int {
	target := 1
	ans := make([]int, num+1)
	for i := 0; i <= num; i++ {
		if i < 2 {
			ans[i] = i
		} else if target*2 == i {
			target *= 2
			ans[i] = 1
		} else {
			ans[i] = ans[i-target] + 1
		}
	}

	return ans
}

func countBitsBest(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
		if i%2 == 0 {
			result[i] = result[i/2]
		} else {
			result[i] = result[i/2] + 1
		}
	}
	return result
}
