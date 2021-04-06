package main

func main() {
	levelOrder(&TreeNode{1, nil, nil})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		var levelAns []int
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i] == nil {
				continue
			}

			levelAns = append(levelAns, queue[i].Val)
			queue = append(queue, queue[i].Left, queue[i].Right)
		}

		if len(levelAns) > 0 {
			ans = append(ans, levelAns)
		}

		queue = queue[l:]
	}
	return ans
}
