package main

import "fmt"

func main() {
	head := &ListNode{Val: 2, Next: nil}
	tail := &ListNode{Val: 0, Next: &ListNode{Val: -4, Next: head}}
	head.Next = tail
	fmt.Println(detectCycle(&ListNode{Val: 3, Next: head}))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast.Next != nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			break
		}
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

//func detectCycle(head *ListNode) *ListNode {
//	fast, slow := head, head
//	for fast != nil && slow != nil {
//		slow = slow.Next
//		fast = fast.Next
//		if fast != nil {
//			fast = fast.Next
//		}
//		if fast == slow {
//			break
//		}
//	}
//
//	if fast == nil || slow == nil {
//		return nil
//	}
//
//	slow = head
//	p := fast
//	for slow != p {
//		p = p.Next
//		if p == fast {
//			slow = slow.Next
//		}
//	}
//
//	return slow
//}
