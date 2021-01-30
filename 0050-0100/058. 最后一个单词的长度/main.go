package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("aaaa haha  "))
}

func lengthOfLastWord(s string) int {
	l, i := 0, len(s)-1
	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		l++
		i--
	}

	return l
}
