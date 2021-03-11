package main

import "fmt"

func main() {
	fmt.Println(calculate("0-2147483647"))
}

// 速度100%, 空间34.90%
func calculate(s string) int {
	var stack []int
	var ops []byte

	for i := 0; i < len(s); i++ {
		l := s[i]
		if l == ' ' {
			continue
		} else if l <= '9' && l >= '0' {
			num := int(l - '0')
			for i+1 < len(s) && s[i+1] <= '9' && s[i+1] >= '0' {
				num = num*10 + int(s[i+1]-'0')
				i++
			}

			if len(ops) == 0 {
				stack = append(stack, num)
				continue
			}

			if ops[len(ops)-1] == '/' {
				stack[len(stack)-1] = stack[len(stack)-1] / num
				ops = ops[:len(ops)-1]
			} else if ops[len(ops)-1] == '*' {
				stack[len(stack)-1] = stack[len(stack)-1] * num
				ops = ops[:len(ops)-1]
			} else {
				stack = append(stack, num)
			}
		} else {
			ops = append(ops, l)
		}
	}

	//fmt.Println(stack, ops)

	for i := 1; i < len(stack); i++ {
		if ops[i-1] == '+' {
			stack[i] += stack[i-1]
		} else {
			stack[i] = stack[i-1] - stack[i]
		}
	}

	return stack[len(stack)-1]
}
