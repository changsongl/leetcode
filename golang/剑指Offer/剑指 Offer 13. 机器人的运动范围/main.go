package main

import "fmt"

func main() {
	fmt.Println(movingCount(16, 8, 4), 15)
	fmt.Println(movingCount(38, 15, 9), 135)
}

func movingCount(m int, n int, k int) int {
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}

	type data struct {
		row int
		col int
	}

	moves := []data{{0, 1}, {1, 0}}
	ans := 1

	for stack := []data{{0, 0}}; len(stack) != 0; stack = stack[1:] {
		d := stack[0]
		for _, move := range moves {
			tx := d.row + move.row
			ty := d.col + move.col
			if tx < 0 || tx >= m || ty < 0 || ty >= n || dp[tx][ty] || get(tx)+get(ty) > k {
				continue
			}
			stack = append(stack, data{tx, ty})
			dp[tx][ty] = true
			ans++
		}
	}

	return ans
}

func get(i int) int {
	v := 0
	for i != 0 {
		v += i % 10
		i /= 10
	}
	return v
}

// my solution alloc over max allowed
//func movingCount(m int, n int, k int) int {
//	dp := make([][]bool, m)
//	for i := 0; i < m; i++ {
//		dp[i] = make([]bool, n)
//	}
//
//	type data struct {
//		row int
//		col int
//	}
//
//	moves := []data{{0, 1}, {1, 0}}
//	ans := 0
//
//	for stack := []data{{0, 0}}; len(stack) != 0; stack = stack[1:] {
//		d := stack[0]
//
//		if get(d.row)+get(d.col) <= k && !dp[d.row][d.col] {
//			dp[d.row][d.col] = true
//			ans++
//		}
//
//		for _, move := range moves {
//			r := d.row + move.row
//			c := d.col + move.col
//			if r >= m || c >= n || dp[r][c] || !dp[d.row][d.col] {
//				continue
//			}
//			stack = append(stack, data{r, c})
//		}
//	}
//
//	return ans
//}
