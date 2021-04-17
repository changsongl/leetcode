package main

import "fmt"

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m := make(map[int]int, k+1)
	for i, n := range nums {
		id := getBucketId(n, t)
		if _, exist := m[id]; exist {
			return true
		} else if large, exists := m[id+1]; exists && large-n <= t {
			return true
		} else if small, exists := m[id-1]; exists && n-small <= t {
			return true
		}

		m[id] = n
		if i >= k {

			delete(m, getBucketId(nums[i-k], t))
		}
	}

	return false
}

func getBucketId(num, t int) int {
	t++
	if num >= 0 {
		return num / t
	}

	return (num+1)/t - 1
}

//func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
//	n := len(nums)
//
//	for i := 0; i < n-1; i++ {
//		for j := i + 1; j < n && j-i <= k; j++ {
//			if abs(nums[i]-nums[j]) <= t || abs(nums[j]-nums[i]) <= t {
//				return true
//			}
//		}
//	}
//
//	return false
//}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
