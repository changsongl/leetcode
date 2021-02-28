package main

import "fmt"

//输入
//[[1,1,0],[1,0,1],[0,0,0]]
//输出
//[[1,1,0],[0,0,0],[1,0,1]]
//预期结果
//[[1,0,0],[0,1,0],[1,1,1]]
func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		l := len(A[i])
		for j := 0; j < l/2 || j == l/2 && l%2 == 1; j++ {
			A[i][j], A[i][l-j-1] = A[i][l-j-1]*-1+1, A[i][j]*-1+1
		}
	}

	return A
}
