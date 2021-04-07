package main

import (
	"fmt"
)

func main() {
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
}

// merge sort solution
func reversePairs(nums []int) int {
	l := len(nums)
	_, count := mergeSort(nums[:l/2], nums[l/2:])
	return count
}

func mergeSort(nums1, nums2 []int) ([]int, int) {
	len1, len2 := len(nums1), len(nums2)
	count := 0
	if len1 > 1 {
		nums1New, countInc := mergeSort(nums1[:len1/2], nums1[len1/2:])
		nums1 = nums1New
		count += countInc
	}

	if len2 > 1 {
		nums2New, countInc := mergeSort(nums2[:len2/2], nums2[len2/2:])
		nums2 = nums2New
		count += countInc
	}

	nums := make([]int, len1+len2)
	i := 0
	for len(nums1) > 0 && len(nums2) > 0 {
		if nums2[0] < nums1[0] {
			nums[i] = nums2[0]
			count += len(nums1)
			nums2 = nums2[1:]
		} else {
			nums[i] = nums1[0]
			nums1 = nums1[1:]
		}
		i++
	}

	if len(nums1) > 0 {
		copy(nums[i:], nums1)
	} else {
		copy(nums[i:], nums2)
	}

	return nums, count
}

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
