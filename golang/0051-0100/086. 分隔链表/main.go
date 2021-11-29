package main

func main() {
	//partition(&ListNode{1, &ListNode{4, &ListNode{3,
	//	&ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}, 3)
	partition(&ListNode{1, &ListNode{1, nil}}, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	small, big := &ListNode{}, &ListNode{}
	smallHead, bigHead := small, big

	for head != nil {
		if head.Val < x {
			small.Next, small = head, head
		} else {
			big.Next, big = head, head
		}
		head = head.Next
	}

	big.Next = nil
	small.Next = bigHead.Next
	return smallHead.Next
}

//func partition(head *ListNode, x int) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	left, right, node := head, head, head.Next
//
//	for node != nil {
//		if node.Val >= x {
//			right.Next, node, right = node, node.Next, node
//		} else if left.Val < x {
//			if right == left {
//				right = node
//			}
//			left, node, left.Next, node.Next = node, node.Next, node, left.Next
//		} else {
//			head = node
//			left, node, node.Next = node, node.Next, left
//		}
//	}
//
//	right.Next = nil
//	return head
//}
