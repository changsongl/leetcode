package main

func main() {
	mergeKLists([]*ListNode{nil, nil})
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	}

	l := mergeKLists(lists[:n/2])
	r := mergeKLists(lists[n/2:])
	return mergeTwoList(l, r)
}

func mergeTwoList(l, r *ListNode) *ListNode {
	dummy := ListNode{0, nil}
	pnt := &dummy
	for l != nil && r != nil {
		if l.Val < r.Val {
			pnt.Next = l
			l = l.Next
		} else {
			pnt.Next = r
			r = r.Next
		}

		pnt = pnt.Next
		pnt.Next = nil
	}

	if l != nil {
		pnt.Next = l
	} else {
		pnt.Next = r
	}

	return dummy.Next
}

//func mergeKLists(lists []*ListNode) *ListNode {
//	if len(lists) == 0 {
//		return nil
//	}
//
//	dummy := ListNode{0, nil}
//	pnt := &dummy
//
//	newList := make([]*ListNode, 0, len(lists))
//	for i := 0; i < len(lists); i++ {
//		if lists[i] != nil {
//			newList = append(newList, lists[i])
//		}
//	}
//	lists = newList
//
//	for len(lists) > 0 {
//		i, v := 0, lists[0].Val
//		for j := 0; j < len(lists); j++ {
//			v2 := lists[j].Val
//			if v2 < v {
//				v = v2
//				i = j
//			}
//		}
//
//		pnt.Next = lists[i]
//		lists[i] = lists[i].Next
//		pnt.Next.Next = nil
//		pnt = pnt.Next
//
//		if lists[i] == nil {
//			lists = append(lists[0:i], lists[i+1:]...)
//		}
//	}
//
//	return dummy.Next
//}
