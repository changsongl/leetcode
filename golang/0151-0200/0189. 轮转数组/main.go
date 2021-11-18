package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6}
	rotate(nums, 4)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func rotateSwap(nums []int, k int) {
	n := len(nums)
	if n < 2 || k%n == 0 {
		return
	}

	num, pos, sPos := nums[0], (0+k)%n, 0

	for i := 0; i < n; i++ {
		if pos != sPos {
			num, nums[pos], pos = nums[pos], num, (pos+k)%n
		} else {
			nums[pos] = num
			sPos++
			num, pos = nums[sPos], (sPos+k)%n
		}
	}
}
