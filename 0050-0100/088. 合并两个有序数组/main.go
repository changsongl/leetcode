package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	index, num1Index, num2Index := m+n-1, m-1, n-1
	for num1Index >= 0 && num2Index >= 0 {
		n1, n2 := nums1[num1Index], nums2[num2Index]
		if n1 > n2 {
			nums1[index] = n1
			num1Index--
		} else {
			nums1[index] = n2
			num2Index--
		}

		index--
	}

	for num2Index >= 0 {
		nums1[index] = nums2[num2Index]
		num2Index--
		index--
	}
}
