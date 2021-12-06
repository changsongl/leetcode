package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	copyMap := make(map[*Node]*Node)
	node := head
	for node != nil {
		copyMap[node] = &Node{
			Val: node.Val,
		}
		node = node.Next
	}

	node = head
	for node != nil {
		copyMap[node].Next = copyMap[node.Next]
		copyMap[node].Random = copyMap[node.Random]
		node = node.Next
	}

	return copyMap[head]
}
