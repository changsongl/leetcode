package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	nums   []int
	origin []int
	size   int
}

func init() {
	rand.Seed(time.Now().Unix())
}

func Constructor(nums []int) Solution {
	l := len(nums)
	return Solution{
		nums:   nums,
		origin: append(make([]int, 0, l), nums...),
		size:   l,
	}
}

func (this *Solution) Reset() []int {
	return this.origin
}

func (this *Solution) Shuffle() []int {
	for i := 0; i < this.size-1; i++ {
		chosen := rand.Intn(this.size-i) + i
		this.nums[chosen], this.nums[i] = this.nums[i], this.nums[chosen]
	}

	return this.nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	obj := Constructor(nums)
	nums1 := obj.Shuffle()
	nums2 := obj.Shuffle()
	fmt.Println(nums1, nums2, obj.Reset())
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
