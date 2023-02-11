package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure1(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}

	q := []*TreeNode{A}

	for len(q) > 0 {
		if isSameStructure(q[0], B) {
			return true
		}

		if q[0] != nil {
			q = append(q, q[0].Left, q[0].Right)
		}

		q = q[1:]
	}

	return false
}

func isSameStructure(A *TreeNode, B *TreeNode) bool {
	return (B == nil) || (A == nil && B == nil) ||
		(A != nil && B != nil && A.Val == B.Val && isSameStructure(A.Left, B.Left) && isSameStructure(A.Right, B.Right))
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (rec(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func rec(A *TreeNode, B *TreeNode) bool {
	return (B == nil) || (A != nil && B != nil && A.Val == B.Val && rec(A.Left, B.Left) && rec(A.Right, B.Right))
}
