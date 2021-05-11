package main

import "fmt"

// 011
// 001

// 001
// 010
// 011

func main() {
	fmt.Println(decode([]int{1, 2}, 1))
}

func decode(encoded []int, first int) []int {
	n := len(encoded) + 1
	ans := make([]int, n)
	ans[0] = first
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] ^ encoded[i-1]
	}

	return ans
}
