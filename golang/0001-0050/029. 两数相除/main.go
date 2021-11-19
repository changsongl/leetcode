package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(-2147483648, 2))
}

func divide(dividend int, divisor int) int {
	return int(divide32(int32(dividend), int32(divisor)))
}

func divide32(dividend int32, divisor int32) int32 {
	isNeg := (dividend ^ divisor) < 0
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}

	if divisor == -1 {
		if dividend == math.MinInt32 && !isNeg {
			return math.MaxInt32
		} else if isNeg {
			return dividend
		} else {
			return -dividend
		}
	}

	var curDivisor, curCount, count int32 = divisor, -1, 0

	for (curDivisor<<1) >= dividend && curDivisor >= (math.MinInt32>>1) && curDivisor != math.MinInt32 {
		curDivisor = curDivisor << 1
		curCount = curCount << 1
	}

	for curDivisor <= divisor && curDivisor != -1 {
		if curDivisor >= dividend {
			dividend -= curDivisor
			if isNeg {
				count += curCount
			} else {
				count -= curCount
			}
		}

		curDivisor = curDivisor >> 1
		curCount = curCount >> 1
	}

	return count
}

//func divide32(dividend int32, divisor int32) int32 {
//	if divisor == math.MinInt32 && dividend != math.MinInt32 {
//		return 0
//	} else if dividend == math.MinInt32 && divisor == -1 {
//		return math.MaxInt32
//	} else if divisor == 1 {
//		return dividend
//	} else if divisor == -1 {
//		return -dividend
//	}
//
//	neg := false
//
//	if dividend > 0 {
//		dividend = -dividend
//		neg = !neg
//	}
//
//	if divisor > 0 {
//		divisor = -divisor
//		neg = !neg
//	}
//
//	var count int32
//	curDivisor := divisor
//	var curCount int32 = 1
//	decr := false
//
//	for curDivisor <= divisor {
//		if dividend <= curDivisor {
//			dividend -= curDivisor
//			count -= curCount
//			if !decr {
//				curCount = curCount << 1
//				curDivisor = curDivisor << 1
//			}
//		} else {
//			decr = true
//			curCount = curCount >> 1
//			if curDivisor != -1 {
//				curDivisor = curDivisor >> 1
//			} else {
//				curDivisor = 0
//			}
//		}
//	}
//
//	if !neg {
//		return -count
//	}
//
//	return count
//}
