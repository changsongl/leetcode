package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	p := dummy
	for p.Next != nil && p.Next.Next != nil {
		if p.Next.Val == p.Next.Next.Val {
			var n *ListNode = nil
			np := p.Next.Next
			for np.Next != nil {
				if np.Next.Val == np.Val {
					np = np.Next
				} else {
					n = np.Next
					break
				}
			}
			p.Next = n
		} else {
			p = p.Next
		}
	}

	return dummy.Next
}
