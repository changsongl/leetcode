package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n := &ListNode{Val: 1,
		Next: &ListNode{Val: 10,
			Next: &ListNode{Val: 5,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 7, Next: nil}}}}}
	n = sortList(n)

	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

func sortList(root *ListNode) *ListNode {
	node := root
	var q []*ListNode
	for node.Next != nil {
		q = append(q, node.Next)
		node.Next = node.Next.Next
		node = node.Next
	}

	// 合并root 和 q
	ans, node := &ListNode{}, root
	pnt := ans
	for node != nil && len(q) != 0 {
		if node.Val < q[len(q)-1].Val {
			pnt.Next = node
			node = node.Next
		} else {
			pnt.Next = q[len(q)-1]
			q = q[:len(q)-1]
		}
		pnt = pnt.Next
	}

	pnt.Next = node
	for len(q) != 0 {
		pnt.Next = q[len(q)-1]
		q = q[:len(q)-1]
		pnt = pnt.Next
	}
	pnt.Next = nil

	return ans.Next
}
