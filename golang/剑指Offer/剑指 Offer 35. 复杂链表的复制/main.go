package main

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		n := cur.Next
		cur.Next = &Node{Val: cur.Val, Random: nil, Next: n}
		cur = n
	}

	cur = head
	for cur != nil {
		if cur.Random == nil {
			cur.Next.Random = nil
		} else {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	cur = head
	newHead := cur.Next
	for cur != nil {
		n := cur.Next.Next
		if n == nil {
			cur.Next = nil
			break
		}
		cur.Next.Next = cur.Next.Next.Next
		cur.Next = n
		cur = cur.Next
	}

	return newHead
}

//func copyRandomList(head *Node) *Node {
//	if head == nil {
//		return nil
//	}
//
//	m := make(map[*Node]*Node, 10)
//	newHead := &Node{Val: head.Val}
//	m[head] = newHead
//
//	pnt := head.Next
//	newPnt := newHead
//	for pnt != nil {
//		newPnt.Next = &Node{Val: pnt.Val}
//		m[pnt] = newPnt.Next
//
//		newPnt = newPnt.Next
//		pnt = pnt.Next
//	}
//
//	pnt = head
//	newPnt = newHead
//	for pnt != nil {
//		if pnt.Random == nil {
//			newPnt.Random = nil
//		} else {
//			newPnt.Random = m[pnt.Random]
//		}
//
//		newPnt = newPnt.Next
//		pnt = pnt.Next
//	}
//
//	return newHead
//}
