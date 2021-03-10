package main

import (
	"container/list"
	"fmt"
)

// my bad implementation
func main() {
	fmt.Println(calculate("2-4-(8+2-6+(8+4-(1)+8-10))"))
}

type Operator byte

const (
	Add      Operator = '+'
	Minus    Operator = '-'
	LeftPren Operator = '('
	RightPre Operator = ')'
	Neg      Operator = '_'
)

func calculate(s string) int {
	nums, ops := list.New(), list.New()
	lastOp := true
	startNum := false

	for i := 0; i < len(s); i++ {
		l := s[i]
		if l == ' ' || (lastOp && l == uint8(Add)) {
			continue
		}

		if lastOp && l == uint8(Minus) {
			ops.PushBack(Neg)
		} else if l == uint8(LeftPren) {
			ops.PushBack(LeftPren)
		} else if l == uint8(RightPre) {
			lastOp = false
			val(nums, ops, false)
		} else if l == uint8(Minus) || l == uint8(Add) {
			b := ops.Back()
			if b != nil {
				op := b.Value.(Operator)
				fmt.Println(string(op))
				if op != LeftPren {
					val(nums, ops, true)
				}
			}
			ops.PushBack(Operator(l))
			lastOp = true
		} else {
			if startNum {
				num1 := nums.Remove(nums.Back()).(int)
				nums.PushBack(num1*10 + int(l-'0'))
			} else {
				nums.PushBack(int(l - '0'))
			}
			startNum = true
			lastOp = false
			continue
		}
		startNum = false
	}

	val(nums, ops, false)
	return nums.Remove(nums.Back()).(int)
}

func val(nums, ops *list.List, once bool) {
	for {
		b := ops.Back()
		if b == nil {
			break
		}
		el := ops.Remove(ops.Back())
		if el == nil {
			break
		}

		op := el.(Operator)
		if op == LeftPren {
			break
		}

		if op == Add {
			num2 := nums.Remove(nums.Back()).(int)
			num1 := nums.Remove(nums.Back()).(int)
			nums.PushBack(num1 + num2)
			//fmt.Printf("%d + %d\n", num1, num2)
		} else if op == Minus {
			num2 := nums.Remove(nums.Back()).(int)
			num1 := nums.Remove(nums.Back()).(int)
			nums.PushBack(num1 - num2)
			//fmt.Printf("%d - %d\n", num1, num2)
		} else if op == Neg {
			num1 := nums.Remove(nums.Back()).(int)
			nums.PushBack(-num1)
			//fmt.Printf("- %d\n d", num1)
		}

		if once {
			break
		}
	}
}
