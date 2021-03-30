package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	current := head
//	next := head.Next
//
//	for next != nil {
//		nn := next.Next
//		next.Next = current
//		current = next
//		next = nn
//	}
//	head.Next = nil
//
//	return current
//}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
