package main

import "fmt"

func main() {
	countAndSay(4)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	output, num := "", countAndSay(n-1)
	for i := 0; i < len(num); i++ {
		digit, count := num[i], 1
		for i+1 < len(num) && digit == num[i+1] {
			i++
			count++
		}
		output += fmt.Sprintf("%d%c", count, digit)
	}
	return output
}
