package main

import "fmt"

func main() {
	fmt.Println(truncateSentence("how are you", 2) + "1111")
	fmt.Println("hahaha"[:6])
}

func truncateSentence(s string, k int) string {
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count++
		}

		if count >= k {
			return s[:i]
		}
	}

	return s
}
