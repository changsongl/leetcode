package main

//执行用时： 8 ms, 在所有 Go 提交中击败了 98.87%
//内存消耗： 7.8 MB, 在所有 Go 提交中击败了 11.44% 的用户

import "math"

type StackElement struct {
	num   int
	front *StackElement
	min   int
}

type MinStack struct {
	s *StackElement
	l int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	ele := &StackElement{num: math.MaxInt, min: math.MaxInt}

	return MinStack{s: ele}
}

func (this *MinStack) Push(x int) {
	this.s = &StackElement{num: x, front: this.s, min: this.s.min}
	if this.s.num < this.s.min {
		this.s.min = this.s.num
	}

	this.l++
}

func (this *MinStack) Pop() {
	if this.l > 0 {
		n := this.s.front
		this.s.front = nil
		this.s = n

		this.l--
	}
}

func (this *MinStack) Top() int {
	return this.s.num
}

func (this *MinStack) Min() int {
	return this.s.min
}
