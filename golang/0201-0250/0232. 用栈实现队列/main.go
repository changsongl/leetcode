package main

import "container/list"

type MyQueue struct {
	stack1 *list.List
	stack2 *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack1.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() != 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	data, ok := this.stack2.Remove(this.stack2.Back()).(int)
	if !ok {
		return -1
	}
	return data
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() != 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}

	ele := this.stack2.Back()
	if ele == nil {
		return -1
	}

	data, ok := ele.Value.(int)
	if !ok {
		return -1
	}
	return data
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack2.Len() == 0 && this.stack1.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
