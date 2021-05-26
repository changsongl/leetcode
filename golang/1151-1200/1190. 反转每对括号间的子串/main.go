package main

import "fmt"

func main() {
	fmt.Println(reverseParentheses("(abcd)"))
	fmt.Println(reverseParentheses("(u(love)i)"))
	fmt.Println(reverseParentheses("(ed(et(oc))el)"))
	fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
}
func reverseParentheses(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			rs, inc := reverseString(s[i+1:], 0)
			i += inc + 1
			ans += rs
		} else {
			ans += string(s[i])
		}
	}
	return ans
}

func reverseString(s string, loop int) (string, int) {
	str := ""
	for i := 0; i < len(s); i++ {
		var newStr string
		if s[i] == '(' {
			rs, ni := reverseString(s[i+1:], loop+1)
			i += ni + 1
			newStr = rs
		} else if s[i] == ')' {
			return str, i
		} else {
			newStr = string(s[i])
		}

		if loop%2 == 0 {
			str = newStr + str
		} else {
			str += newStr
		}
	}
	return str, 0
}
