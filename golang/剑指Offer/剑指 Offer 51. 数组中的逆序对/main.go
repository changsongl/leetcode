package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reversePairs([]int{7, 5, 5, 5, 6, 4}))
}

// TODO: 树状数组解法
func reversePairs(nums []int) int {
	n := len(nums)
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)
	fmt.Println(tmp)

	for i := 0; i < n; i++ {
		nums[i] = sort.SearchInts(tmp, nums[i]) + 1
	}
	fmt.Println(tmp, nums)

	bit := BIT{
		n:    n,
		tree: make([]int, n+1),
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		q := bit.query(nums[i] - 1)
		ans += q
		fmt.Printf("%v\n", bit)
		bit.update(nums[i])

		fmt.Printf("%d - %v\n", q, bit)
	}
	return ans
}

type BIT struct {
	n    int
	tree []int
}

func (b BIT) lowbit(x int) int { return x & (-x) }

func (b BIT) query(x int) int {
	ret := 0
	for x > 0 {
		ret += b.tree[x]
		x -= b.lowbit(x)
	}
	return ret
}

func (b BIT) update(x int) {
	for x <= b.n {
		b.tree[x]++
		x += b.lowbit(x)
	}
}

//func reversePairs(nums []int) int {
//	return mergeSort(nums, 0, len(nums)-1)
//}

//func mergeSort(nums []int, start, end int) int {
//	if start >= end {
//		return 0
//	}
//	mid := start + (end-start)/2
//	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
//	tmp := []int{}
//	i, j := start, mid+1
//	for i <= mid && j <= end {
//		if nums[i] <= nums[j] {
//			tmp = append(tmp, nums[i])
//			cnt += j - (mid + 1)
//			i++
//		} else {
//			tmp = append(tmp, nums[j])
//			j++
//		}
//	}
//	for ; i <= mid; i++ {
//		tmp = append(tmp, nums[i])
//		cnt += end - (mid + 1) + 1
//	}
//	for ; j <= end; j++ {
//		tmp = append(tmp, nums[j])
//	}
//	for i := start; i <= end; i++ {
//		nums[i] = tmp[i-start]
//	}
//	return cnt
//}

// merge sort solution
//func reversePairs(nums []int) int {
//	l := len(nums)
//	_, count := mergeSort(nums[:l/2], nums[l/2:])
//	return count
//}
//
//func mergeSort(nums1, nums2 []int) ([]int, int) {
//	len1, len2 := len(nums1), len(nums2)
//	count := 0
//	if len1 > 1 {
//		nums1New, countInc := mergeSort(nums1[:len1/2], nums1[len1/2:])
//		nums1 = nums1New
//		count += countInc
//	}
//
//	if len2 > 1 {
//		nums2New, countInc := mergeSort(nums2[:len2/2], nums2[len2/2:])
//		nums2 = nums2New
//		count += countInc
//	}
//
//	nums := make([]int, len1+len2)
//	i := 0
//	for len(nums1) > 0 && len(nums2) > 0 {
//		if nums2[0] < nums1[0] {
//			nums[i] = nums2[0]
//			count += len(nums1)
//			nums2 = nums2[1:]
//		} else {
//			nums[i] = nums1[0]
//			nums1 = nums1[1:]
//		}
//		i++
//	}
//
//	if len(nums1) > 0 {
//		copy(nums[i:], nums1)
//	} else {
//		copy(nums[i:], nums2)
//	}
//
//	return nums, count
//}

// brute force
//func reversePairs(nums []int) int {
//	count := 0
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i] > nums[j] {
//				count++
//			}
//		}
//	}
//	return count
//}

// bubble sort (over time)
//func reversePairs(nums []int) int {
//	count := 0
//	for i := 0; i < len(nums)-1; i++ {
//		for j := 0; j < len(nums)-i-1; j++ {
//			if nums[j] > nums[j+1] {
//				nums[j], nums[j+1] = nums[j+1], nums[j]
//				count++
//			}
//		}
//	}
//	return count
//}
