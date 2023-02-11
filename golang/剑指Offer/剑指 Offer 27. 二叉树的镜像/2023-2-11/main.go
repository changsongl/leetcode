package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return &TreeNode{Val: root.Val, Left: mirrorTree(root.Right), Right: mirrorTree(root.Left)}
}
