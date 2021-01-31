package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}

func subsets(nums []int) [][]int {
	var result [][]int

	for _, num := range nums {
		if len(result) == 0 {
			result = append(result, []int{}, []int{num})
			continue
		}

		for _, set := range result {
			l := len(set)
			var newSet = make([]int, l+1)
			copy(newSet, set)
			newSet[l] = num

			result = append(result, newSet)
		}
	}

	return result
}
