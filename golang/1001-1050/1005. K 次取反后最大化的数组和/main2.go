package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
}

func largestSumAfterKNegations(nums []int, k int) int {
	h := Heap(nums)

	heap.Init(&h)

	for k > 0 {
		num := heap.Pop(&h).(int)
		if num >= 0 {
			if k%2 == 0 {
				num = -num
			}
			k = 0
		}

		heap.Push(&h, -num)
		k--
	}

	num := 0
	for h.Len() != 0 {
		num += heap.Pop(&h).(int)
	}
	return num
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	n := h.Len()
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}
