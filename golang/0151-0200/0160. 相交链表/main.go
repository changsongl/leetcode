package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Other Solution相交链表，获得两个链表长度，长的先走N步，然后短的再走。当相遇时则代表是交点。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nA, nB := headA, headB
	for nA != nB {
		if nA == nil {
			nA = headB
		} else {
			nA = nA.Next
		}
		if nB == nil {
			nB = headA
		} else {
			nB = nB.Next
		}
	}

	return nA
}
