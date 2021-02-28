package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针，快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

// map 去重
func hasCycleMap(head *ListNode) bool {
	m := map[*ListNode]bool{}
	for head != nil {
		if _, exist := m[head]; exist {
			return true
		}

		m[head] = true
		head = head.Next
	}

	return false
}
