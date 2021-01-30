package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n := rotateRight(
		&ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5, Next: nil}}}}}, 2,
	)

	fmt.Println(n.Val)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head, _ = rotateRightSub(head, head, k, 0)
	return head
}

func rotateRightSub(head, current *ListNode, k, c int) (*ListNode, int) {
	c++

	if current.Next == nil {
		if k%c == 0 {
			return head, -1
		}

		p := c - k%c + 1
		current.Next = head
		return current, p
	}

	h, p := rotateRightSub(head, current.Next, k, c)
	if c == p {
		return current, p
	} else if c == p-1 {
		current.Next = nil
	}

	return h, p
}
