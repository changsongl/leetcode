package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 转换成双链表交接问题
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent, visited := make(map[*TreeNode]*TreeNode), make(map[*TreeNode]bool)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			parent[node.Left] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			dfs(node.Right)
		}
	}

	dfs(root)

	for q != nil {
		visited[q] = true
		q = parent[q]
	}

	for p != nil {
		if visited[p] {
			return p
		}
		p = parent[p]
	}

	return nil
}

//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	if root == nil || root == p || root == q {
//		return root
//	}
//
//	left := lowestCommonAncestor(root.Left, p, q)
//	right := lowestCommonAncestor(root.Right, p, q)
//	if left != nil && right != nil {
//		return root
//	} else if left != nil {
//		return left
//	} else {
//		return right
//	}
//}
