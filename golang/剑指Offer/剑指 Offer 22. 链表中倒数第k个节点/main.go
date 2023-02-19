package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd1(head *ListNode, k int) *ListNode {
	var s []*ListNode

	for head != nil {
		s = append(s, head)
		head = head.Next
	}

	return s[len(s)-k]
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ {
		tail = tail.Next
	}

	for tail != nil {
		head = head.Next
		tail = tail.Next
	}

	return head
}
