package main

import "fmt"

func main() {
	fmt.Println(removeNthFromEnd(
		&ListNode{1, nil}, 1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var cur = 0

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeNthFromEnd(head.Next, n)
	cur++
	if cur == n {
		return head.Next
	}
	return head
}

// 快慢指针
func removeNthFromEndTwoPointer(head *ListNode, n int) *ListNode {
	prev, current := head, head
	for i := 0; i < n-1; i++ {
		current = current.Next
	}

	if current.Next == nil {
		return prev.Next
	}

	current = current.Next
	for current.Next != nil {
		prev = prev.Next
		current = current.Next
	}

	prev.Next = prev.Next.Next

	return head
}
