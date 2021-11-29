package main

import (
	"sort"
)

func main() {
	kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3)
}

//输入：arr = [1,2,3,5], k = 3
//输出：[2,5]
//解释：已构造好的分数,排序后如下所示:
//1/5, 1/3, 2/5, 1/2, 3/5, 2/3
//很明显第三个最小的分数是 2/5

type PrimeFractions [][]int

func (p PrimeFractions) Less(i, j int) bool {
	return p[i][0]*p[j][1] < p[i][1]*p[j][0]
}

func (p PrimeFractions) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PrimeFractions) Len() int {
	return len(p)
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)

	result := make([][]int, 0, n*(n-1))
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			result = append(result, []int{arr[i], arr[j]})
		}
	}

	sort.Sort(PrimeFractions(result))
	return result[k-1]
}
