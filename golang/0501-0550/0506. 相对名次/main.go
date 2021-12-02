package main

import (
	"container/heap"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	n := len(score)
	if n == 0 {
		return []string{}
	}

	m := map[int]string{1: "Gold Medal", 2: "Silver Medal", 3: "Bronze Medal"}
	ans := make([]string, n)
	indexMap := make(map[int]int, n)
	for i, s := range score {
		indexMap[s] = i
	}

	h := Heap(score)
	heap.Init(&h)

	i := 1

	for h.Len() != 0 {
		num := heap.Pop(&h).(int)
		if i > 3 {
			ans[indexMap[num]] = strconv.FormatInt(int64(i), 10)
		} else {
			ans[indexMap[num]] = m[i]
		}
		i++
	}

	return ans
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
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
