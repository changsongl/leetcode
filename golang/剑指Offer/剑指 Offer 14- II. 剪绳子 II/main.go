package main

import "fmt"

func main() {
	fmt.Println(cuttingRopeMath(120))
}

// 数学解法当 n1+n2+...nn = n, Max(n1*n2*...nn)，分段最接近3
func cuttingRopeMath(n int) int {
	if n <= 3 {
		return n - 1
	}

	a, b := n/3, n%3
	if b == 0 {
		for i := 0; i < a; i++ {

		}
		return pow(a)
	} else if b == 1 {
		return (pow(a-1) * 4) % 1000000007
	} else {
		return (pow(a) * 2) % 1000000007
	}
}

func pow(y int) int {
	if y == 0 {
		return 1
	}

	ans := 1
	for i := 0; i < y; i++ {
		ans = (ans * 3) % 1000000007
	}
	return ans
}
