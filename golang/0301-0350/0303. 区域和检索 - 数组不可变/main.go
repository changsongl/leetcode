package main

type NumArray struct {
	cheatBook []int
}

func Constructor(nums []int) NumArray {
	l := len(nums)
	na := NumArray{}
	na.cheatBook = make([]int, l)
	if l != 0 {
		na.cheatBook[0] = nums[0]

		for i := 1; i < l; i++ {
			na.cheatBook[i] = na.cheatBook[i-1] + nums[i]
		}
	}

	return na
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.cheatBook[j]
	} else {
		return this.cheatBook[j] - this.cheatBook[i-1]
	}
}
