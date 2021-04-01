package main

import "fmt"

func main() {
	fmt.Println(clumsy(10))
}

// * / + -
func clumsy(N int) int {
	if N <= 2 {
		return N
	}

	num := N
	nums := num * (num - 1) / (num - 2)
	num -= 3

	for num >= 4 {
		nums += num - (num-1)*(num-2)/(num-3)
		num -= 4
	}

	if num != 0 {
		nums += 1
	}

	return nums
}
