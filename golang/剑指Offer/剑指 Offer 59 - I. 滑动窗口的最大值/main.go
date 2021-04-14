package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 1))
}

func maxSlidingWindow(nums []int, k int) []int {
	hp := NewHeap(len(nums), nums)
	heap.Init(&hp)
	n := len(nums)
	ans := make([]int, 0, n-k)
	for i := 0; i < n; i++ {
		heap.Push(&hp, i)

		if i >= k-1 {
			for len(hp.S) > 0 && hp.S[0] < i-k+1 {
				heap.Pop(&hp)
			}
			ans = append(ans, hp.Top())
		}
	}

	return ans
}

func NewHeap(l int, q []int) IntHeap {
	return IntHeap{S: make([]int, 0, l), Q: q}
}

// An IntHeap is a min-heap of ints.
type IntHeap struct {
	S []int
	Q []int
}

func (h IntHeap) Len() int           { return len(h.S) }
func (h IntHeap) Less(i, j int) bool { return h.Q[h.S[i]] > h.Q[h.S[j]] }
func (h IntHeap) Swap(i, j int)      { h.S[i], h.S[j] = h.S[j], h.S[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.S = append(h.S, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := len(h.S)
	x := h.S[n-1]
	h.S = h.S[:n-1]
	return h.Q[x]
}

func (h *IntHeap) Top() int {
	x := h.S[0]
	return h.Q[x]
}
