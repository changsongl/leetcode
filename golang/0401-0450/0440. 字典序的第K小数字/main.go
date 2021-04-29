package main

import "fmt"

func main() {
	fmt.Println(findKthNumber(13, 2))
}

func findKthNumber(n int, k int) int {
	var getCount func(prefix int) int
	getCount = func(prefix int) int {
		cur, next := prefix, prefix+1
		count := 0

		for cur <= n {
			count += min(next, n+1) - cur
			cur *= 10
			next *= 10
		}

		return count
	}

	num := 1
	count := 1
	for count < k {
		curCount := getCount(num)
		if count+curCount <= k {
			num++
			count += curCount
		} else {
			count++
			num *= 10
		}
	}

	return num
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func findKthNumberBruteForce(n int, k int) int {
//	count := 0
//	var dfs func(num int) bool
//	ans := 0
//
//	dfs = func(num int) bool {
//		count++
//		if count == k {
//			ans = num
//			return true
//		}
//
//		if num == 1 || num%10 == 0 {
//			for i := 0; i < 10; i++ {
//				newNum := num
//				if i == 0 {
//					newNum *= 10
//				} else {
//					newNum += i
//				}
//				if newNum <= n {
//					if result := dfs(newNum); result {
//						return true
//					}
//				}
//			}
//		} else {
//			newNum := num * 10
//			if newNum <= n {
//				if result := dfs(newNum); result {
//					return true
//				}
//			}
//		}
//		return false
//	}
//
//	dfs(1)
//	return ans
//}
