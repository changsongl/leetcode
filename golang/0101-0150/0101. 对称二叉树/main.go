package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricSubNodes(root.Left, root.Right)
}

func isSymmetricSubNodes(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	}

	result := isSymmetricSubNodes(left.Left, right.Right)
	if !result {
		return false
	}

	return isSymmetricSubNodes(left.Right, right.Left)
}
