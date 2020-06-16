package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(findMedianSortedArraysEasy([]int{1000001}, []int{1000000}))
}

func findMedianSortedArraysEasy(nums1 []int, nums2 []int) float64 {
	var nums []int
	i1, i2, l1, l2 := 0, 0, len(nums1), len(nums2)
	if l1 == 0 {
		nums = append(nums, nums2[:]...)
	}else if l2 == 0{
		nums = append(nums, nums1[:]...)
	}
	for{
		if i1 == l1{
			nums = append(nums, nums2[i2:]...)
			break
		}else if i2 == l2{
			nums = append(nums, nums1[i1:]...)
			break
		}

		if nums1[i1] > nums2[i2]{
			nums = append(nums, nums2[i2])
			i2++
		}else{
			nums = append(nums, nums1[i1])
			i1++
		}
	}

	l := l1 + l2
	if (l1 + l2) % 2 == 0{
		l--
	}
	mid := int(math.Ceil(float64(l)/2)) - 1
	if (l1 + l2) % 2 == 0{
		return float64(nums[mid] + nums[mid+1])/2
	}else{
		return float64(nums[mid])
	}
}