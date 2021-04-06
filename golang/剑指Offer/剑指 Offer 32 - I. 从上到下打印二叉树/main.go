package main

import "fmt"

func main() {
	fmt.Println(levelOrder(&TreeNode{0, nil, nil}))
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	queue := []*TreeNode{root}
	var ans []int

	for len(queue) != 0 {
		node := queue[0]
		if node != nil {
			ans = append(ans, node.Val)
			queue = append(queue, node.Left, node.Right)
		}

		queue = queue[1:]
	}
	return ans
}
