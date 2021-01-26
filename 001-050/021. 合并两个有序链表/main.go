package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// recursion
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
}

// loop to merge
func mergeTwoListsFor(l1 *ListNode, l2 *ListNode) *ListNode {
	var orderedList = &ListNode{}
	current := orderedList

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			current.Next = l2
			l2 = l2.Next
		} else {
			current.Next = l1
			l1 = l1.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else if l2 != nil {
		current.Next = l2
	}

	return orderedList.Next
}
