package main

import "fmt"

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5},
		[]int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}

func maxSatisfied(customers []int, grumpy []int, X int) int {
	happyNum := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			happyNum += customers[i]
			customers[i] = 0
		}
	}

	count, grumpyMax, grumpyCurrent := X, 0, 0
	for i := 0; i < len(customers); i++ {
		if count == 0 {
			grumpyCurrent -= customers[i-X]
		} else {
			count--
		}

		grumpyCurrent += customers[i]
		grumpyMax = max(grumpyMax, grumpyCurrent)
	}

	return happyNum + grumpyMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
