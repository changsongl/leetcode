package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	prev, ans := -1, math.MaxInt64
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if prev != -1 && root.Val-prev < ans {
			ans = root.Val - prev
		}
		prev = root.Val
		dfs(root.Right)
	}

	dfs(root)
	return ans
}

//func minDiffInBST(root *TreeNode) int {
//	s := BFS(root)
//	n := len(s)
//	if n == 0 {
//		return -1
//	}
//
//	diff := math.MaxInt64
//	for i := 1; i < n; i++ {
//		diff = int(math.Min(math.Abs(float64(s[i]-s[i-1])), float64(diff)))
//	}
//
//	return diff
//}
//
//func BFS(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	var tree []int
//	tree = append(tree, BFS(root.Left)...)
//	tree = append(tree, root.Val)
//	tree = append(tree, BFS(root.Right)...)
//
//	return tree
//}
