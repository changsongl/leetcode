package main

import "fmt"

func main() {
	fmt.Println(integerReplacement(8))
}

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	} else if n == 3 {
		return 2
	} else if n%2 == 0 {
		return integerReplacement(n/2) + 1
	} else if n%4 == 3 {
		return integerReplacement(n+1) + 1
	} else if n%4 == 1 {
		return integerReplacement(n-1) + 1
	}

	return 0
}

//func integerReplacement(n int) int {
//	var dp = map[int]int{1: 0, 2: 1, 3: 2}
//
//	var f func(num int) int
//	f = func(num int) int {
//		if num == 1 {
//			return 0
//		} else if dp[num] != 0 {
//			return dp[num]
//		}
//
//		var count int
//
//		if num%2 == 0 {
//			count = f(num/2) + 1
//		} else {
//			count = min(f(num+1), f(num-1)) + 1
//		}
//
//		dp[num] = count
//		return count
//	}
//
//	return f(n)
//}
//
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//
//	return a
//}
