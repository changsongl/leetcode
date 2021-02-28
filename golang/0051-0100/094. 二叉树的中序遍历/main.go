package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris
func inorderTraversalMorris(root *TreeNode) []int {
	for root != nil {
		if root.Left != nil {
			// predecessor 节点就是当前 root 节点向左走一步，然后一直向右走至无法走为止

			// 让 predecessor 的右指针指向 root，继续遍历左子树 if
			// 说明左子树已经访问完了，我们需要断开链接 else
		} else { // 如果没有左孩子，则直接访问右孩子

		}

	}
	return nil
}

// recursion
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	l := append(inorderTraversal(root.Left), root.Val)
	r := inorderTraversal(root.Right)

	return append(l, r...)
}
