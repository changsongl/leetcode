package main

import "fmt"

// slow
//func judgeSquareSum(c int) bool {
//	dict := make(map[int]bool)
//	for i := 0; i <= c; i++ {
//		if i*i > c {
//			break
//		}
//		dict[i*i] = true
//	}
//
//	for num, _ := range dict {
//		if dict[c-num] {
//			return true
//		}
//	}
//	return false
//}

func main() {
	fmt.Println(judgeSquareSum(999999999))
}

func judgeSquareSum(c int) bool {
	// 优化r的值
	l, r := 0, 0
	for r*r <= c {
		r++
	}

	for l <= r {
		num := l*l + r*r
		if num == c {
			return true
		} else if num > c {
			r--
		} else {
			l++
		}
	}
	return false
}

// 超时
//func judgeSquareSum(c int) bool {
//	l, r := 0, c
//
//	for l <= r {
//		num := l*l + r*r
//		if num == c {
//			return true
//		} else if num > c {
//			r--
//		} else {
//			l++
//		}
//	}
//	return false
//}
