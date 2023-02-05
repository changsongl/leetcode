package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint1(head *ListNode) []int {
	nums := make([]int, 0, 10000)
	l := 0
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
		l++
	}

	for i := 0; i < l/2; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}

	return nums
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	return append(reversePrint(head.Next), head.Val)
}
