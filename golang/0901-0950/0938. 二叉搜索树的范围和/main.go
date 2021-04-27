package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		ele := queue[0]
		if ele != nil {
			if isMiddle(ele.Val, low, high) {
				sum += ele.Val
			}
			if ele.Val >= low && ele.Left != nil {
				queue = append(queue, ele.Left)
			}

			if ele.Val <= high && ele.Right != nil {
				queue = append(queue, ele.Right)
			}
		}
		queue = queue[1:]
	}

	return sum
}

func isMiddle(num, low, high int) bool {
	return num >= low && num <= high
}
