package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	increasingBST(&TreeNode{3, &TreeNode{Val: 1}, &TreeNode{Val: 6}})
}

func increasingBST(root *TreeNode) *TreeNode {
	newRoot := &TreeNode{}
	last := newRoot

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		node.Left = nil
		last.Right = node
		last = node
		inorder(node.Right)
	}

	inorder(root)
	return newRoot.Right
}
