package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := maxDepth(root.Left)
	if right := maxDepth(root.Right); right > max {
		max = right
	}
	return max + 1
}
