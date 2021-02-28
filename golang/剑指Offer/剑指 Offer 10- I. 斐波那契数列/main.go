package main

import "fmt"

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
}

var fibSlice [101]int

func fib(n int) int {
	return Fib(n) % (1e9 + 7)
}

func Fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	fib2 := fibSlice[n-2]
	if fib2 == 0 {
		fib2 = fib(n - 2)
		fibSlice[n-2] = fib2
	}

	fib1 := fibSlice[n-1]
	if fib1 == 0 {
		fib1 = fib(n - 1)
		fibSlice[n-1] = fib1
	}

	return fib2 + fib1
}
