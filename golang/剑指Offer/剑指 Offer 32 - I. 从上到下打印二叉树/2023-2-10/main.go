package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var ans []int

	nodes := []*TreeNode{root}

	for len(nodes) != 0 {
		if nodes[0] != nil {
			ans = append(ans, nodes[0].Val)
			nodes = append(nodes, nodes[0].Left, nodes[0].Right)
		}

		nodes = nodes[1:]
	}

	return ans
}
