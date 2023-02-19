package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	var ans string
	l, r := 0, 0

	for r < len(s) {
		if s[l] == ' ' && s[r] == ' ' {
			l++
			r++
		} else if s[r] == ' ' {
			ans = fmt.Sprintf("%s %s", s[l:r], ans)
			r++
			l = r
		} else {
			r++
		}
	}

	if l < len(s) && s[l] != ' ' {
		ans = fmt.Sprintf("%s %s", s[l:r], ans)
	}

	return strings.TrimRight(ans, " ")
}
