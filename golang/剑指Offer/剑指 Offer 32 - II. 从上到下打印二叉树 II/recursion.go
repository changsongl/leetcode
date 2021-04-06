package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func levelOrder(root *TreeNode) [][]int {
//	var ans [][]int
//	return levelOrderCore(root, 0, ans)
//}
//
//func levelOrderCore(node *TreeNode, level int, ans [][]int) [][]int {
//	if node != nil {
//		if len(ans) == level {
//			ans = append(ans, []int{})
//		}
//		ans[level] = append(ans[level], node.Val)
//
//		ans = levelOrderCore(node.Left, level+1, ans)
//		ans = levelOrderCore(node.Right, level+1, ans)
//	}
//
//	return ans
//}
