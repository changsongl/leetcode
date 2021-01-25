package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("{()})"))
}

// slow
func isValid(s string) bool {
	l := len(s)
	for l != 0 {
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "()", "", -1)

		newL := len(s)
		if newL == l {
			return false
		}
		l = newL
	}

	return true
}

// stack
func isValidStack(s string) bool {
	var braceMap = map[int32]int32{
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := make([]int32, 0, len(s)/2)

	for _, c := range s {
		openBrace, isClose := braceMap[c]
		if !isClose {
			stack = append(stack, c)
			continue
		}

		lStack := len(stack)
		if len(stack) == 0 {
			return false
		}

		openBraceInStack := stack[lStack-1]
		if openBraceInStack != openBrace {
			return false
		}
		stack = stack[:lStack-1]
	}

	return len(stack) == 0
}
