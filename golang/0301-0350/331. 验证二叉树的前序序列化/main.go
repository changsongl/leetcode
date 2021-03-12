package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	fmt.Println(isValidSerialization("1,#"))
}

func isValidSerialization(preorder string) bool {
	if preorder == "#" {
		return true
	}

	nodes := strings.Split(preorder, ",")
	count := len(nodes)
	for i, node := range nodes {
		if node == "#" {
			replaces := replace(i, nodes)
			if replaces == -1 {
				return false
			}

			count = count - replaces - 1
		}
	}

	return count == 0
}

func replace(i int, nodes []string) int {
	for j := i - 1; j >= 0; j-- {
		if nodes[j] == "#" {
			continue
		}

		if nodes[j] == "*" {
			nodes[j] = "#"
			if j == 0 {
				return 1
			}

			next := replace(j, nodes)
			if next == -1 {
				return -1
			}

			return next + 1
		} else {
			nodes[j] = "*"
			return 0
		}
	}
	return -1
}
