package main

import "fmt"

func main() {
	fmt.Println(decode([]int{3, 1}))
}

func decode(encoded []int) []int {
	total := 0
	n := len(encoded)
	for i := 1; i <= n+1; i++ {
		total ^= i
	}

	a := 0
	for i := 1; i < n; i += 2 {
		a ^= encoded[i]
	}

	ans := make([]int, n+1)
	ans[0] = total ^ a
	for i := 0; i < n; i++ {
		ans[i+1] = encoded[i] ^ ans[i]
	}

	return ans
}
