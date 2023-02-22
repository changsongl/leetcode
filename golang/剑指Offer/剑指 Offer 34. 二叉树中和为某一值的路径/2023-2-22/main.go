package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeTotal struct {
	node  *TreeNode
	total int
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}

	var ans [][]int

	parentMap := make(map[*TreeNode]*TreeNode)
	pathFunc := func(leaf *TreeNode) (path []int) {
		for leaf != nil {
			path = append(path, leaf.Val)
			leaf = parentMap[leaf]
		}

		l := len(path)
		for i := 0; i < l/2; i++ {
			path[i], path[l-i-1] = path[l-i-1], path[i]
		}

		return
	}

	queue := []*NodeTotal{{node: root, total: root.Val}}
	for len(queue) != 0 {
		node := queue[0]
		if node.total == target && node.node.Left == nil && node.node.Right == nil {
			ans = append(ans, pathFunc(node.node))
		}

		if node.node.Left != nil {
			parentMap[node.node.Left] = node.node
			queue = append(queue, &NodeTotal{node: node.node.Left, total: node.total + node.node.Left.Val})
		}

		if node.node.Right != nil {
			parentMap[node.node.Right] = node.node
			queue = append(queue, &NodeTotal{node: node.node.Right, total: node.total + node.node.Right.Val})
		}

		queue = queue[1:]
	}

	return ans
}
