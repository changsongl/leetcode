package main

func main() {
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	node1, node2 := headA, headB
//	for node1 != node2 {
//		if node1 != nil {
//			node1 = node1.Next
//		} else {
//			node1 = headA
//		}
//
//		if node2 != nil {
//			node2 = node2.Next
//		} else {
//			node2 = headB
//		}
//
//	}
//	return node1
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	node1, node2 := headB, headA

	for node1 != node2 {
		if node1 != nil {
			node1 = node1.Next
		} else {
			node1 = headA
		}

		if node2 != nil {
			node2 = node2.Next
		} else {
			node2 = headB
		}
	}

	return node1
}

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	dict := make(map[*ListNode]bool)
//	for headA != nil {
//		dict[headA] = true
//		headA = headA.Next
//	}
//
//	for headB != nil {
//		_, exists := dict[headB]
//		if exists {
//			return headB
//		}
//		headB = headB.Next
//	}
//
//	return nil
//}
