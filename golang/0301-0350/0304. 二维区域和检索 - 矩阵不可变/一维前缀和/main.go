package main

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))
	for i, row := range matrix {
		sums[i] = make([]int, len(row)+1)
		for j, v := range row {
			sums[i][j+1] = sums[i][j] + v
		}
	}
	return NumMatrix{sums}
}

func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) (sum int) {
	for i := row1; i <= row2; i++ {
		sum += nm.sums[i][col2+1] - nm.sums[i][col1]
	}
	return
}
