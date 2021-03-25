package main

func main() {
	swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}})
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := ListNode{0, head}
	pnt := &dummy
	for pnt.Next != nil && pnt.Next.Next != nil {
		one := pnt.Next
		two := pnt.Next.Next

		one.Next = two.Next
		two.Next = one
		pnt.Next = two
		pnt = one
	}

	return dummy.Next
}
