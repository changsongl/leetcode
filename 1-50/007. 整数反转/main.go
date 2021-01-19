package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	r := 0
	for x > 9 || x < -9 {
		r = r*10 + x%10
		x /= 10
	}

	if r > 0 && (r > math.MaxInt32/10 || r == math.MaxInt32/10 && x > math.MaxInt32%10) {
		return 0
	} else if r < 0 && (r < math.MinInt32/10 || r == math.MinInt32/10 && x < math.MinInt32%10) {
		return 0
	}

	return r*10 + x%10
}

func reverseSmart(x int) int {
	var r int32 = 0

	for x != 0 {
		r *= 10
		if r%10 != 0 {
			return 0
		}
		r += int32(x % 10)
		x /= 10
	}

	return int(r)
}

func reverseInt64(x int) int {
	var r int64 = 0

	for x != 0 {
		r = r*10 + int64(x%10)
		x /= 10
	}

	if r > math.MaxInt32 || r < math.MinInt32 {
		return 0
	}
	return int(r)
}

func reverseWorst(x int) int {
	isNeg := false
	s := strconv.Itoa(x)
	rev := make([]byte, 0, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			isNeg = true
			continue
		}
		rev = append(rev, s[i])
	}

	num, err := strconv.ParseInt(string(rev), 10, 64)
	if err != nil {
		panic(err)
	}

	if isNeg {
		num = -num
	}

	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}

	return int(num)
}
