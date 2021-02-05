package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSwap([]int{0, 7, 8, 10, 10, 11, 12, 13, 19, 18}, []int{4, 4, 5, 7, 11, 14, 15, 16, 17, 20}))
}

func minSwap(A []int, B []int) int {
	n1, s1 := 0, 1
	for i := 1; i < len(A); i++ {
		n2, s2 := math.MaxInt32, math.MaxInt32
		if A[i-1] < A[i] && B[i-1] < B[i] {
			if n2 > n1 {
				n2 = n1
			}
			if s2 > s1+1 {
				s2 = s1 + 1
			}
		}

		if A[i-1] < B[i] && B[i-1] < A[i] {
			if n2 > s1 {
				n2 = s1
			}
			if s2 > n1+1 {
				s2 = n1 + 1
			}
		}

		n1 = n2
		s1 = s2
	}

	if n1 > s1 {
		return s1
	}

	return n1
}
