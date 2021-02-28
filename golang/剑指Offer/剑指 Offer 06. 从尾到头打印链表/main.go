package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrintRecursive(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	s := reversePrintRecursive(head.Next)
	return append(s, head.Val)
}

func reversePrint(head *ListNode) []int {
	s := make([]int, 0, 10)
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}

	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}

	return s
}
