package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {

	tilt := 0

	var rec func(node *TreeNode) int
	rec = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		vLeft, vRight := rec(node.Left), rec(node.Right)

		cTilt := vLeft - vRight
		if cTilt < 0 {
			cTilt = -cTilt
		}
		tilt += cTilt

		return node.Val + vLeft + vRight
	}

	rec(root)

	return tilt
}
