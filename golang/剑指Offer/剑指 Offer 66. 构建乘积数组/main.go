package main

import (
	"fmt"
)

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
	//fmt.Println(constructArr([]int{}))
	//fmt.Println(constructArr([]int{1}))
	//fmt.Println(constructArr([]int{1, 2}))
}

func constructArr(a []int) []int {
	if a == nil || len(a) == 0 {
		return nil
	}
	n := len(a)
	b := make([]int, n)

	left := 1
	for i := 0; i < n; i++ {
		b[i] = left
		left = left * a[i]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		b[i] = b[i] * right
		right = right * a[i]
	}

	return b
}

//func constructArr(a []int) []int {
//	n := len(a)
//	if n == 1 {
//		return []int{0}
//	}
//
//	left, right := make([]int, n), make([]int, n)
//
//	for i := 0; i < n-1; i++ {
//		if i == 0 {
//			left[0] = a[0]
//		} else {
//			left[i] = a[i] * left[i-1]
//		}
//	}
//	for i := n - 1; i > 0; i-- {
//		if i == n-1 {
//			right[n-1] = a[n-1]
//		} else {
//			right[i] = a[i] * right[i+1]
//		}
//	}
//
//	ans := make([]int, n)
//	for i := 0; i < n; i++ {
//		if i == 0 {
//			ans[i] = right[i+1]
//		} else if i == n-1 {
//			ans[i] = left[n-2]
//		} else {
//			ans[i] = left[i-1] * right[i+1]
//		}
//	}
//
//	return ans
//}

//func constructArr(a []int) []int {
//	prod := 0
//	zero := 0
//	for _, num := range a {
//		if num != 0 {
//			if prod == 0 {
//				prod = num
//			} else {
//				prod *= num
//			}
//		} else {
//			zero++
//		}
//	}
//
//	ans := make([]int, 0, len(a))
//
//	for _, num := range a {
//		if num == 0 {
//			if zero > 1 {
//				ans = append(ans, 0)
//			} else {
//				ans = append(ans, prod)
//			}
//		} else {
//			if zero > 0 {
//				ans = append(ans, 0)
//			} else {
//				ans = append(ans, int(float64(prod)*math.Pow(float64(num), -1)))
//			}
//		}
//	}
//
//	return ans
//}
