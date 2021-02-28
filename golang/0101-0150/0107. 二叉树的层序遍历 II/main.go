package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
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

	l := len(levelOrder)
	for i := 0; i < l/2; i++ {
		levelOrder[i], levelOrder[l-i-1] = levelOrder[l-i-1], levelOrder[i]
	}
	return levelOrder
}

//
//func levelOrderBottom(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//	return levelOrders([]*TreeNode{root})
//}
//
//func levelOrders(nodes []*TreeNode) [][]int {
//	var result [][]int
//	var nodeSlice []*TreeNode
//	var currentLevel []int
//	for _, node := range nodes {
//		if node.Left != nil {
//			nodeSlice = append(nodeSlice, node.Left)
//		}
//		if node.Right != nil {
//			nodeSlice = append(nodeSlice, node.Right)
//		}
//
//		currentLevel = append(currentLevel, node.Val)
//	}
//
//	if len(nodes) > 0 {
//		result = append(result, levelOrders(nodeSlice)...)
//	}
//
//	if len(currentLevel) > 0 {
//		result = append(result, currentLevel)
//	}
//
//	return result
//}
