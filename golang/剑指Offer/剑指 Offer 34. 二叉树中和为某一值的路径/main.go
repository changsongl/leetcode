package main

import "fmt"

func main() {
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(pathSum(&root, 1))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	var path []int
	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Right == nil && root.Left == nil {
			if target-root.Val == 0 {
				ans = append(ans, append([]int(nil), path...))
			}
		}

		dfs(root.Left, target-root.Val)
		dfs(root.Right, target-root.Val)
		path = path[:len(path)-1]
	}

	dfs(root, target)
	return ans
}
