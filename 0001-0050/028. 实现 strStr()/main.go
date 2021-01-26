package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))
}

// TODO: KMP 优化
func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	for i := 0; i <= len(haystack)-needleLen; i++ {
		j := 0

		for j < needleLen && haystack[i+j] == needle[j] {
			j++
		}

		if j == needleLen {
			return i
		}
	}
	return -1
}
