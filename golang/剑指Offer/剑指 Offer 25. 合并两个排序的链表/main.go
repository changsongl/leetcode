package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	header := &ListNode{}
	cur := header

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next, cur, l2 = l2, l2, l2.Next
		} else {
			cur.Next, cur, l1 = l1, l1, l1.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return header.Next
}
