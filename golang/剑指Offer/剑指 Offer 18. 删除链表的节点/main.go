package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	} else if head.Val == val {
		return head.Next
	}

	node := head
	for node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			break
		}
		node = node.Next
	}
	return head
}
