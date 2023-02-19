package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	header := &ListNode{Next: head}
	prev, next := header, header.Next

	for next != nil {
		if next.Val == val {
			prev.Next = next.Next
			break
		}

		prev, next = prev.Next, next.Next
	}

	return header.Next
}
