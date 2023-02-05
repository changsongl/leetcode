package main

// Node Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList1(head *Node) *Node {
	cur := head
	m := make(map[*Node]*Node)
	for cur != nil {
		m[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		m[cur].Next, m[cur].Random = m[cur.Next], m[cur.Random]
		cur = cur.Next
	}

	return m[head]
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
	}

	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}

	newHead := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		newNode, curNext := cur.Next, cur.Next.Next
		if curNext != nil {
			newNode.Next = curNext.Next
		}
		cur.Next = curNext
	}

	return newHead
}
