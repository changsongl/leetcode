package main

import "fmt"

func main() {
	fmt.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}), true)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		result []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}

//func canFinish(numCourses int, prerequisites [][]int) bool {
//	graph := make(map[int][]int)
//	nodes := make(map[int]bool)
//
//	for _, p := range prerequisites {
//		graph[p[0]] = append(graph[p[0]], p[1])
//	}
//
//	var traversal func(node int, dup map[int]bool) bool
//	traversal = func(node int, dup map[int]bool) bool {
//		if dup[node] {
//			return false
//		} else if nodes[node] {
//			return true
//		}
//
//		dup[node] = true
//		nodes[node] = true
//		for _, next := range graph[node] {
//			if !traversal(next, dup) {
//				return false
//			}
//		}
//		dup[node] = false
//		return true
//	}
//
//	for node := range graph {
//		if !traversal(node, map[int]bool{}) {
//			return false
//		}
//	}
//
//	return numCourses >= len(graph)
//}
