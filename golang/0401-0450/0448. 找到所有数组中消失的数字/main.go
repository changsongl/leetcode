package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

//执行用时：56 ms, 在所有 Go 提交中击败了90.57%的用户
//内存消耗：7.4 MB, 在所有 Go 提交中击败了72.03%的用户
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num == i+1 {
			continue
		}

		for true {
			nextNum := nums[num-1]
			if nextNum == num {
				break
			}

			nums[num-1] = num
			num = nextNum
		}
	}

	var ans []int

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ans = append(ans, i+1)
		}
	}

	return ans
}
