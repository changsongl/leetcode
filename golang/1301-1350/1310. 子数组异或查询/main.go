package main

import "fmt"

func main() {
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
}

func xorQueries(arr []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i := 1; i < len(arr); i++ {
		arr[i] ^= arr[i-1]
	}

	for i, q := range queries {
		ans[i] = arr[q[1]]
		if q[0]-1 >= 0 {
			ans[i] ^= arr[q[0]-1]
		}
	}

	return ans
}
