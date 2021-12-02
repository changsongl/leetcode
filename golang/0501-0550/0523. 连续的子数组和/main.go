package main

import "fmt"

func main() {
	//fmt.Println(checkSubarraySum([]int{1, 2, 3}, 5))
	fmt.Println(checkSubarraySum([]int{23, 0, 0}, 6))
}

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	m := make(map[int]int, n)
	m[nums[0]%k] = 0
	for i := 1; i < n; i++ {
		nums[i] = (nums[i] + nums[i-1]) % k
		if nums[i] == 0 {
			return true
		}

		index, exist := m[nums[i]]
		if exist {
			if i-index > 1 {
				return true
			}
		} else {
			m[nums[i]] = i
		}
	}
	return false
}

func checkSubarraySumN2(nums []int, k int) bool {
	for i := range nums {
		pre := 0
		for j := i; j < len(nums); j++ {
			pre += nums[j]
			if pre%k == 0 && j > i {
				return true
			}
		}
	}

	return false
}
