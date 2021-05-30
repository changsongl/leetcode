package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(0))
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(2))
	fmt.Println(isPowerOfTwo(17))
}

//func isPowerOfTwo(n int) bool {
//	if n <= 0 {
//		return false
//	}
//	countOne := 0
//	for n != 0 {
//		if n%2 == 1 {
//			countOne++
//		}
//		n = n >> 1
//	}
//	return countOne == 1
//}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
