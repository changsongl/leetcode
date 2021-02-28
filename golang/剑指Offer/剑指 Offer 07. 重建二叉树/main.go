package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
// [ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
//
// 不断找根节点。
func buildTreeFindRoot(preorder []int, inorder []int) *TreeNode {
	l := len(preorder)
	if l == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	i := 0
	for i < l && inorder[i] != root.Val {
		i++
	}

	root.Left = buildTreeFindRoot(preorder[1:i+1], inorder[:i])
	root.Right = buildTreeFindRoot(preorder[i+1:], inorder[i+1:])

	return root
}

// 巧妙解法
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{root}
	inorderIndex := 0

	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}
