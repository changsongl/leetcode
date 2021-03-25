package main

func main() {
	reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3,
		&ListNode{4, &ListNode{5, nil}}}}}, 3)
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	pnt := head
	for c := 0; c < k; c++ {
		if pnt == nil {
			return head
		}
		pnt = pnt.Next
	}

	r := pnt
	first := head
	last := head
	for c := 0; c < k; c++ {
		second := first.Next
		first.Next = r
		r = first
		first = second
	}

	last.Next = reverseKGroup(last.Next, k)

	return r
}
