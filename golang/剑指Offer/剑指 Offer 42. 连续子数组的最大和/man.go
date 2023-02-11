package main

type status struct {
	iSum, lSum, rSum, mSum int
}

func maxSubArray(nums []int) int {
	return getMax(nums, 0, len(nums)-1).mSum
}

func getMax(nums []int, l, r int) status {
	if l == r {
		return status{nums[l], nums[l], nums[l], nums[l]}
	}

	m := (r + l) / 2
	return merge(getMax(nums, l, m), getMax(nums, m+1, r))
}

func merge(l, r status) status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return status{iSum, lSum, rSum, mSum}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxSubArray1(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}

		if max < nums[i] {
			max = nums[i]
		}
	}

	return max
}
