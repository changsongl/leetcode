package main

import "fmt"

func main() {
	fmt.Println(myPow(-1, 2147483647))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	ans := x

	for i := 1; i < n; i++ {
		ans *= x
		if ans == 0 {
			return 0
		} else if ans == 1 {
			if x == -1 && (n-i-1)%2 == 1 {
				return -1
			}
			return 1
		}
	}

	return ans
}
