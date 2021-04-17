package main

import "math"

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	leftX := int(math.Max(float64(A), float64(E)))
	leftY := int(math.Max(float64(B), float64(F)))
	rightX := int(math.Min(float64(C), float64(G)))
	rightY := int(math.Min(float64(D), float64(H)))

	total := (C-A)*(D-B) + (G-E)*(H-F)
	if D <= F || A >= G || C <= E || B >= H {
		return total
	}
	return total - (rightX-leftX)*(rightY-leftY)
}

//func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
//	if A > E {
//		return computeArea(E, F, G, H, A, B, C, D)
//	}
//
//	// 调整两个矩形位置, 让第一个矩形靠最左边
//	if A > E {
//		return computeArea(E, F, G, H, A, B, C, D)
//	}
//	// 没有重叠的情况
//	if B >= H || D <= F || C <= E {
//		return abs(A-C)*abs(B-D) + abs(E-G)*abs(F-H)
//	}
//
//	// 重叠情况
//	// 下边界
//	down := max(A, E)
//	// 上
//	up := min(C, G)
//	// 左
//	left := max(B, F)
//	// 右
//	right := min(D, H)
//	return abs(A-C)*abs(B-D) + abs(E-G)*abs(F-H) - abs(up-down)*abs(left-right)
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}
//
//func abs(a int) int {
//	if a < 0 {
//		return -a
//	}
//	return a
//}
//
////输入：
////-2 A
////-2 B
////2 C
////2 D
////-2 E
////-2 F
////2 G
////2 H
