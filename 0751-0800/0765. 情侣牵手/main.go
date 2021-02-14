package main

import "fmt"

func main() {
	fmt.Println(minSwapsCouples([]int{0, 2, 1, 3}))
}

// greedy
func minSwapsCouples(row []int) int {
	l := len(row)
	seatMap := make([]int, l)
	for i := 0; i < l; i++ {
		seatMap[row[i]] = i
	}

	c := 0

	for i := 0; i < l-1; i += 2 {
		left := row[i]
		right := left + 1
		if left%2 == 1 {
			right = left - 1
		}

		rPos := seatMap[right]
		if rPos != i+1 {
			row[rPos] = row[i+1]
			seatMap[row[i+1]] = rPos
			c++
		}
	}

	return c
}
