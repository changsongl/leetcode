package main

//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//
//	var res [][]int
//	var list1, list2 []*TreeNode
//	list1 = append(list1, root)
//
//	for len(list1) > 0 || len(list2) > 0 {
//		if len(list1) > 0 {
//			var subRes []int
//			for _, node := range list1 {
//				subRes = append(subRes, node.Val)
//				if node.Left != nil {
//					list2 = append(list2, node.Left)
//				}
//				if node.Right != nil {
//					list2 = append(list2, node.Right)
//				}
//			}
//			res = append(res, subRes)
//			list1 = list1[:0]
//		}
//
//		if len(list2) > 0 {
//			var subRes []int
//			for _, node := range list2 {
//				subRes = append(subRes, node.Val)
//				if node.Left != nil {
//					list1 = append(list1, node.Left)
//				}
//				if node.Right != nil {
//					list1 = append(list1, node.Right)
//				}
//			}
//			res = append(res, subRes)
//			list2 = list2[:0]
//		}
//	}
//	return res
//
//}
