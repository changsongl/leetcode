package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || isSymmetricSub(root.Left, root.Right)
}

func isSymmetricSub(a, b *TreeNode) bool {
	return (a == nil && b == nil) || (a != nil && b != nil &&
		a.Val == b.Val && isSymmetricSub(a.Left, b.Right) && isSymmetricSub(a.Right, b.Left))
}

// use queue push left and right, and pop two nodes for each check
