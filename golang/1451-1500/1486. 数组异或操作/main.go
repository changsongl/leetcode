package main

import "fmt"

func main() {
	fmt.Println(xorOperation(4, 3))
}

//func xorOperation(n int, start int) int {
//	ans := 0
//	for i := 0; i < n; i++ {
//		ans ^= start + i*2
//	}
//	return ans
//}

//func xorOperation(n int, start int) int {
//	ans := 0
//	for i := 1; i < n; i++ {
//		ans += 1
//		ans <<= 1
//	}
//	return ans ^ start
//}

func sumXor(x int) int {
	switch x % 4 {
	case 0:
		return x
	case 1:
		return 1
	case 2:
		return x + 1
	default:
		return 0
	}
}

// 数学法
func xorOperation(n, start int) (ans int) {
	s, e := start>>1, n&start&1
	ret := sumXor(s-1) ^ sumXor(s+n-1)
	return ret<<1 | e
}
