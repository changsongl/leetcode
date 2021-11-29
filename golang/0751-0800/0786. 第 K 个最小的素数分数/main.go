package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 7, 11, 12}, 4))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	h := make(hp, n-1)
	for j := 1; j < n; j++ {
		h[j-1] = frac{arr[0], arr[j], 0, j}
	}
	heap.Init(&h)
	for loop := k - 1; loop > 0; loop-- {
		f := heap.Pop(&h).(frac)
		fmt.Println(f)
		if f.i+1 < f.j {
			heap.Push(&h, frac{arr[f.i+1], f.y, f.i + 1, f.j})
		}
	}
	return []int{h[0].x, h[0].y}
}

type frac struct{ x, y, i, j int }
type hp []frac

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].x*h[j].y < h[i].y*h[j].x }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(frac)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

//输入：arr = [1,2,3,5], k = 3
//输出：[2,5]
//解释：已构造好的分数,排序后如下所示:
//1/5, 1/3, 2/5, 1/2, 3/5, 2/3
//很明显第三个最小的分数是 2/5

//type PrimeFractions [][]int
//
//func (p PrimeFractions) Less(i, j int) bool {
//	return p[i][0]*p[j][1] < p[i][1]*p[j][0]
//}
//
//func (p PrimeFractions) Swap(i, j int) {
//	p[i], p[j] = p[j], p[i]
//}
//
//func (p PrimeFractions) Len() int {
//	return len(p)
//}
//
//func kthSmallestPrimeFraction(arr []int, k int) []int {
//	n := len(arr)
//
//	result := make([][]int, 0, n*(n-1))
//	for i := 0; i < n; i++ {
//		for j := i + 1; j < n; j++ {
//			result = append(result, []int{arr[i], arr[j]})
//		}
//	}
//
//	sort.Sort(PrimeFractions(result))
//	return result[k-1]
//}
