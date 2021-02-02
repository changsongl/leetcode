package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(2))
}

func mySqrt(x int) int {
	// 二分查找
	l, r := 0, x

	for l <= r {
		m := (l + r) / 2
		sqr := m * m

		if sqr == x {
			return m
		} else if sqr > x {
			r = m - 1
		} else {
			l = l + 1
		}
	}

	if r*r > x {
		return r - 1
	} else {
		return r
	}
}

func newTownRule(x int) int {
	if x < 0 {
		panic("negative number")
	} else if x == 0 {
		return 0
	}

	var cur, fx = 1.0, float64(x)
	for {
		pre := cur
		cur = (cur + fx/cur) / 2
		if math.Abs(cur-pre) < 1e-6 {
			return int(cur)
		}
	}
}
