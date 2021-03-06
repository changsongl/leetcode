package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

// 分组，prefix 和 suffix 通过对比 两部分 拿出区间内最高值
//func maxSlidingWindow(nums []int, k int) []int {
//	n := len(nums)
//	prefixMax := make([]int, n)
//	suffixMax := make([]int, n)
//	for i, v := range nums {
//		if i%k == 0 {
//			prefixMax[i] = v
//		} else {
//			prefixMax[i] = max(prefixMax[i-1], v)
//		}
//	}
//
//	for i := n - 1; i >= 0; i-- {
//		if i == n-1 || (i+1)%k == 0 {
//			suffixMax[i] = nums[i]
//		} else {
//			suffixMax[i] = max(suffixMax[i+1], nums[i])
//		}
//	}
//
//	fmt.Println(prefixMax, suffixMax)
//
//	ans := make([]int, n-k+1)
//	for i := range ans {
//		ans[i] = max(suffixMax[i], prefixMax[i+k-1])
//	}
//	return ans
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func maxSlidingWindow(nums []int, k int) []int {
	hp := NewHeap(len(nums), nums)
	heap.Init(&hp)
	n := len(nums)
	ans := make([]int, 0, n-k)
	for i := 0; i < n; i++ {

		if i >= k-1 {
			num := nums[i]
			for len(hp.S) > 0 && (hp.S[0] < i-k+1 || num > hp.Top()) {
				heap.Pop(&hp)
			}
			heap.Push(&hp, i)
			ans = append(ans, hp.Top())
		} else {
			heap.Push(&hp, i)
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
