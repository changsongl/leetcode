package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB

	for pA != pB {
		if pA == nil {
			pA = headA
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headB
		} else {
			pB = pB.Next
		}
	}

	return pA
}
