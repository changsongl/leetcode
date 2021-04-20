package main

import "fmt"

func main() {
	fmt.Println(strStr("", ""))
}

func strStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pmt := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pmt[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pmt[i] = j
	}

	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pmt[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

//func strStr(haystack, needle string) int {
//	n, m := len(haystack), len(needle)
//	if m == 0 {
//		return 0
//	}
//	pi := make([]int, m)
//	for i, j := 1, 0; i < m; i++ {
//		fmt.Printf("around: i = %d, j = %d\n", i, j)
//		for j > 0 && needle[i] != needle[j] {
//			j = pi[j-1]
//			fmt.Printf("j = %d\n", j)
//		}
//		if needle[i] == needle[j] {
//			j++
//			fmt.Printf("j ++ = %d\n", j)
//		}
//		pi[i] = j
//	}
//	fmt.Println(pi)
//	for i, j := 0, 0; i < n; i++ {
//		for j > 0 && haystack[i] != needle[j] {
//			j = pi[j-1]
//		}
//		if haystack[i] == needle[j] {
//			j++
//		}
//		if j == m {
//			return i - m + 1
//		}
//	}
//	return -1
//}

//func strStr(haystack string, needle string) int {
//	needleLen := len(needle)
//	for i := 0; i <= len(haystack)-needleLen; i++ {
//		j := 0
//
//		for j < needleLen && haystack[i+j] == needle[j] {
//			j++
//		}
//
//		if j == needleLen {
//			return i
//		}
//	}
//	return -1
//}
