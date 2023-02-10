package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	nodes := []*TreeNode{root}

	for len(nodes) != 0 {
		var levelNodes []*TreeNode
		var levelAns []int

		for _, n := range nodes {
			if n == nil {
				continue
			}

			levelAns = append(levelAns, n.Val)
			levelNodes = append(levelNodes, n.Left, n.Right)
		}

		if len(levelAns) != 0 {
			ans = append(ans, levelAns)
		}

		nodes = levelNodes
	}

	return ans
}
