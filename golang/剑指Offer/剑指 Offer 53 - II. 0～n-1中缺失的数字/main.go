package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3, 4, 5, 6, 7, 8, 9}))
	//fmt.Println(missingNumber([]int{0, 1, 3}))
	//fmt.Println(missingNumber([]int{0, 2, 3}))
}

func missingNumber(nums []int) int {
	l, r := 0, len(nums)
	for r > l {
		mid := l + (r-l)>>1
		fmt.Println(l, r, mid)
		if nums[mid] == mid {
			if l == mid {
				break
			}
			l = mid
		} else {
			r = mid
		}
	}
	return r

}

//func missingNumber(nums []int) int {
//	n := len(nums)
//	l, r := 0, n-1
//
//	for l <= r {
//		m := (l + r) / 2
//		if nums[m] == m {
//			l = m + 1
//		} else {
//			r = m - 1
//		}
//	}
//
//	return -1
//}

//func missingNumber(nums []int) int {
//	n := len(nums)
//	l, r := 0, n-1
//
//	for l <= r {
//		m := l + (r-l)>>1
//		if nums[m] == m {
//			l = m + 1
//			if l >= n {
//				return n
//			} else if nums[l] != l {
//				return l
//			}
//		} else {
//			r = m - 1
//			if r < 0 {
//				return 0
//			} else if nums[r] == r {
//				return r + 1
//			}
//		}
//	}
//
//	return -1
//}
