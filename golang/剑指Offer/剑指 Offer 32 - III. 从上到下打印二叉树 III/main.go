package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//             1
//       2          3
//    4   5       6   7
//  8  9 10 11  12 13 14 15

// 1
// 3 2
// 4 5 6 7
// 15 14 13 12 11 10 9 8

func levelOrder1(root *TreeNode) [][]int {
	var ans [][]int
	nodes := []*TreeNode{root}
	var level int

	for len(nodes) != 0 {
		var levelNodes []*TreeNode
		var levelAns []int

		for _, n := range nodes {
			if n == nil {
				continue
			}

			if level%2 != 0 {
				levelAns = append([]int{n.Val}, levelAns...)
			} else {
				levelAns = append(levelAns, n.Val)
			}

			levelNodes = append(levelNodes, n.Left, n.Right)
		}

		if len(levelAns) != 0 {
			ans = append(ans, levelAns)
		}

		level++
		nodes = levelNodes
	}

	return ans
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	nodes := []*TreeNode{root}
	var level int
	var levelAns []int

	for len(nodes) != 0 {
		nodes, levelAns = getLevelAnswers(level, nodes)
		if len(levelAns) != 0 {
			ans = append(ans, levelAns)
		}

		level++
	}

	return ans
}

func getLevelAnswers(level int, nodes []*TreeNode) ([]*TreeNode, []int) {
	var levelNodes []*TreeNode
	var levelAns []int

	for i := 0; i < len(nodes); i++ {
		n := nodes[i]
		if n == nil {
			continue
		}

		if level%2 != 0 {
			levelAns = append([]int{n.Val}, levelAns...)
		} else {
			levelAns = append(levelAns, n.Val)
		}
		levelNodes = append(levelNodes, n.Left, n.Right)
	}

	return levelNodes, levelAns
}
