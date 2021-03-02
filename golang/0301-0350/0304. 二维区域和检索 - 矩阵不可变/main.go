package main

import "fmt"

func main() {
	nm := Constructor([][]int{{1, 2, 3}, {4, 5, 6}})
	nm.print()
	fmt.Println(nm.SumRegion(1, 1, 1, 2))
}

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	nm := NumMatrix{
		dp: make([][]int, len(matrix)),
	}

	for row := 0; row < len(matrix); row++ {
		rowSlice := make([]int, len(matrix[row]))
		nm.dp[row] = rowSlice
		for col := 0; col < len(matrix[row]); col++ {
			v := matrix[row][col]
			if row > 0 {
				v += nm.dp[row-1][col]
			}
			if col > 0 {
				v += nm.dp[row][col-1]
			}
			if row > 0 && col > 0 {
				v -= nm.dp[row-1][col-1]
			}
			rowSlice[col] = v
		}
	}
	return nm
}

func (this *NumMatrix) print() {
	for row := 0; row < len(this.dp); row++ {
		for col := 0; col < len(this.dp[row]); col++ {
			fmt.Printf("%d, ", this.dp[row][col])
		}
		fmt.Println()
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.dp[row2][col2]
	} else if row1 > 0 && col1 > 0 {
		return this.dp[row2][col2] - this.dp[row1-1][col2] - this.dp[row2][col1-1] + this.dp[row1-1][col1-1]
	} else if col1 == 0 {
		return this.dp[row2][col2] - this.dp[row1-1][col2]
	} else {
		return this.dp[row2][col2] - this.dp[row2][col1-1]
	}
}
