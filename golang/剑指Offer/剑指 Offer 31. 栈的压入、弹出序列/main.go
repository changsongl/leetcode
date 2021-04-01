package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	popLen, pushLen := len(popped), len(pushed)
	if pushLen > popLen {
		return false
	}

	stack := make([]int, 0, pushLen)
	j := 0
	for i := 0; i <= pushLen && j < popLen; i++ {
		if i != pushLen {
			stack = append(stack, pushed[i])
		}
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}

	return j == popLen
}
