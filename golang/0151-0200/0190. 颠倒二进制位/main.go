package main

import "fmt"

func main() {
	fmt.Println(reverseBits(43261596))
}

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 | num&1
		num >>= 1
	}
	return res
}

//var dict []uint32
//
//func init() {
//	dict = make([]uint32, 32)
//	dict[0] = 1
//	for i := 1; i < 32; i++ {
//		dict[i] = 2 * dict[i-1]
//	}
//}
//
//func reverseBits(num uint32) uint32 {
//	i := 31
//	var ans uint32 = 0
//	for num != 0 && i >= 0 {
//		if num&1 == 1 {
//			ans += dict[i]
//		}
//		num /= 2
//		i--
//	}
//	return ans
//}
