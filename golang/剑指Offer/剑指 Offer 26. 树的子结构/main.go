package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 待优化
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}

	same := isSameStructure(A, B)
	return same || (A != nil && (isSubStructure(A.Left, B) || isSubStructure(A.Right, B)))
}

func isSameStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	} else if A == nil {
		return false
	} else if A.Val != B.Val {
		return false
	}

	if ok := isSameStructure(A.Left, B.Left); !ok {
		return false
	} else if ok := isSameStructure(A.Right, B.Right); !ok {
		return false
	}

	return true
}
