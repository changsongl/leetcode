package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func searchBST(root *TreeNode, val int) *TreeNode {
//	if root == nil {
//		return nil
//	} else if root.Val == val {
//		return root
//	}
//
//	if root.Val < val {
//		return searchBST(root.Right, val)
//	} else {
//		return searchBST(root.Left, val)
//	}
//}

func searchBST(root *TreeNode, val int) *TreeNode {
	node := root

	for node != nil {
		if node.Val == val {
			return node
		} else if node.Val > val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return nil
}
