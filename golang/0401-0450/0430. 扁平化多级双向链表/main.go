package main

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node1.Next = node2
	node2.Prev = node1
	node2.Next = node3
	node3.Prev = node2
	node4.Next = node5
	node5.Prev = node4
	node2.Child = node4

	flatten(node1)
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	var f func(n *Node) *Node
	f = func(n *Node) *Node {
		last, node := n, n

		for node != nil {
			last = node
			node = node.Next

			if last.Child != nil {
				nextLast := f(last.Child)
				last.Next, last.Child.Prev = last.Child, last
				if node != nil {
					node.Prev = nextLast
					nextLast.Next = node
				}
				last.Child = nil
			}
		}

		return last
	}

	f(root)

	return root
}
