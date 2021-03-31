package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	fmt.Println(minStack.Top(), minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Top(), minStack.Min())
}

type MinStack struct {
	main []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.main = append(this.main, x)
	if len(this.min) == 0 || x <= this.Min() {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	if this.Top() == this.Min() {
		this.min = this.min[:len(this.min)-1]
	}
	this.main = this.main[:len(this.main)-1]
}

func (this *MinStack) Top() int {
	return this.main[len(this.main)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
