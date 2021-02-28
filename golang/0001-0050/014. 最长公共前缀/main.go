package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefixShort([]string{"flower", "flow", "flight"}))
}

// less code, but time complexity is m*n*m
func longestCommonPrefixShort(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}

	str := strs[0]
	for i := 1; i < l; i++ {
		for strings.Index(strs[i], str) != 0 {
			str = str[0 : len(str)-1]
			fmt.Println(str)
		}
	}

	return str
}

// more code with time complexity n*m
func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 1 {
		return strs[0]
	}

	charIndex := 0
	var prefix []uint8
	var checkChar uint8 = 0

	for i := 0; i < l; i = (i + 1) % l {
		s := strs[i]

		if charIndex >= len(s) {
			break
		} else if i == 0 {
			checkChar = s[charIndex]
		} else if checkChar != s[charIndex] {
			break
		}

		if i == l-1 {
			prefix = append(prefix, checkChar)
			charIndex++
		}
	}

	return string(prefix)
}
