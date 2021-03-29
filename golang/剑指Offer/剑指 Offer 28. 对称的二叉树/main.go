package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricCore(root.Left, root.Right)
}

func isSymmetricCore(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else if node1.Val == node2.Val {
		return isSymmetricCore(node1.Left, node2.Right) && isSymmetricCore(node1.Right, node2.Left)
	} else {
		return false
	}
}
