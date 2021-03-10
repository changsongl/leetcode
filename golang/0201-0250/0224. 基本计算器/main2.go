package main

func calculate2(s string) int {
	stack := []int{}
	length := len(s)
	sign, res := 1, 0
	for i := 0; i < length; i++ {
		ch := s[i]
		if '0' <= ch && ch <= '9' {
			cur := int(ch - '0')
			for i+1 < length && '0' <= s[i+1] && s[i+1] <= '9' {
				cur = cur*10 + int(s[i+1]-'0')
				i++
			}
			res = res + sign*cur
		} else if ch == '+' {
			sign = 1
		} else if ch == '-' {
			sign = -1
		} else if ch == '(' {
			stack = append(stack, res)
			res = 0
			stack = append(stack, sign)
			sign = 1
		} else if ch == ')' {
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			res = num1*res + num2
			stack = stack[:len(stack)-2]
		}
	}
	return res
}
