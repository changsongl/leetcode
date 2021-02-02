package main

func main() {
	//[-2,1,-3,4,-1,2,1,-5,4]
	//fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

//func maxSubArray(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//
//	mid := len(nums) / 2
//	leftMax := maxSubArray(nums[0:mid])
//	rightMax := maxSubArray(nums[mid:])
//
//}

func maxSubArrayOne(nums []int) int {
	result := nums[0]
	sum := 0

	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}

		if sum > result {
			result = sum
		}
	}

	return result
}

// two loop
func maxSubArrayDouble(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]

	for i := 0; i < len(nums); i++ {
		v := 0
		for j := i; j < len(nums); j++ {
			v += nums[j]
			if v > result {
				result = v
			}
		}
	}

	return result
}
