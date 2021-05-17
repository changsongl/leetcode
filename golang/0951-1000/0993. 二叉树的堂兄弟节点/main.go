package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	xFound, yFound := false, false
	xDept, yDept := 0, 0
	var xParent, yParent *TreeNode

	var dfs func(node, parent *TreeNode, dept int)
	dfs = func(node, parent *TreeNode, dept int) {
		if node == nil {
			return
		} else if node.Val == x {
			xFound = true
			xParent = parent
			xDept = dept
			return
		} else if node.Val == y {
			yFound = true
			yParent = parent
			yDept = dept
			return
		}

		if xFound && yFound {
			return
		}

		dfs(node.Left, node, dept+1)
		dfs(node.Right, node, dept+1)
	}

	dfs(root, nil, 0)
	return xDept == yDept && xParent != yParent
}
