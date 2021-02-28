package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var levelOrder [][]int
	if root == nil {
		return nil
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		l := len(queue)
		level := make([]int, 0, l)

		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelOrder = append(levelOrder, level)
	}
	return levelOrder
}
