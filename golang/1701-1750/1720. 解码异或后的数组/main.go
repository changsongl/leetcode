package main

import "fmt"

// 1
// 2

// 001
// 010

// 001
// 000
// 101

// 001
// 000
// 010

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
