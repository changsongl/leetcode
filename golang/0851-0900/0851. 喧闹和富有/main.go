package main

import (
	"fmt"
)

//输入：richer = [[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]], quiet = [3,2,5,4,6,1,7,0]
//输出：[5,5,2,5,4,5,6,7]

func main() {
	fmt.Println(loudAndRich([][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}))
}

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	g := make([][]int, n)
	indeg := make([]int, n)
	q := make([]int, 0, n)
	ans := make([]int, n)
	for _, r := range richer {
		g[r[0]] = append(g[r[0]], r[1])
		indeg[r[1]]++
	}

	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
		ans[i] = i
	}

	for len(q) != 0 {
		for _, node := range g[q[0]] {
			if quiet[ans[q[0]]] < quiet[ans[node]] {
				ans[node] = ans[q[0]]
			}
			indeg[node]--
			if indeg[node] == 0 {
				q = append(q, node)
			}
		}
		q = q[1:]
	}

	return ans
}
