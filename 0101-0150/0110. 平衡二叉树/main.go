package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if balance(root) == -1 {
		return false
	}

	return true
}

func balance(n *TreeNode) int {
	if n == nil {
		return 0
	}

	l := balance(n.Left)
	r := balance(n.Right)
	if l == -1 || r == -1 || (l-r > 1 || l-r < -1) {
		return -1
	}

	l++
	r++

	if l > r {
		return l
	} else {
		return r
	}
}
