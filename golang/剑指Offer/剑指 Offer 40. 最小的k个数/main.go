package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
}

func getLeastNumbersSort(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func getLeastNumbers(arr []int, k int) []int {
	hp := NewHeap(len(arr))
	heap.Init(&hp)
	for i := 0; i < len(arr); i++ {
		heap.Push(&hp, arr[i])
	}

	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(&hp).(int))
	}
	return ans
}

func NewHeap(l int) IntHeap {
	h := make([]int, 0, l)
	return h
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
