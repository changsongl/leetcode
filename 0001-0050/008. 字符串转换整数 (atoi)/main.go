package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi(""))
}

// my solution
func myAtoi(s string) int {
	l := len(s)
	var i = 0
	for i < l && s[i] == ' ' {
		i++
	}
	if l == i {
		return 0
	}

	neg := false
	if s[i] == '-' {
		neg = true
		i++
	} else if s[i] == '+' {
		i++
	}

	var num int32 = 0
	for ; i < l && s[i] <= '9' && s[i] >= '0'; i++ {
		digit := int32(s[i] - '0')
		if neg {
			digit = -digit
		}

		newNum := num*10 + digit
		if newNum/10 != num {
			num = math.MaxInt32
			if neg {
				num = math.MinInt32
			}
			break
		}
		num = newNum
	}

	return int(num)
}
