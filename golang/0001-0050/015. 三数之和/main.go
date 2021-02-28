package main

import (
	"fmt"
	"math"
	"sort"
)

//[-1,0,1,2,-1,-4,-2,-3,3,0,4]

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}

// 44 ms
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	length := len(nums)
	for i := 0; i < length-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r, target := i+1, length-1, -nums[i]
		for l != r {
			moveL, moveR := false, false
			lNum, rNum := nums[l], nums[r]
			if lNum+rNum == target {
				result = append(result, []int{nums[i], lNum, rNum})
				moveL, moveR = true, true
			} else {
				if nums[i]+lNum+rNum > 0 {
					moveR = true
				} else {
					moveL = true
				}
			}

			for moveL && l < r {
				l++
				if nums[l] != lNum {
					break
				}
			}

			for moveR && l < r {
				r--
				if nums[r] != rNum {
					break
				}
			}
		}
	}

	return result
}

// 492 ms	10 MB
func threeSumSort(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var result [][]int
	last := nums[0] - 1
	for i := 0; i < len(nums)-2; i++ {
		if last == nums[i] {
			continue
		}
		last = nums[i]

		target := -nums[i]
		m := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			m[nums[j]] += 1
		}

		for j := i + 1; j < len(nums); j++ {
			if m[target-nums[j]] == 0 {
				continue
			}

			if nums[j]*2 == target {
				n, t := m[nums[j]], 0
				for n = n - 1; n > 0; n-- {
					t += n
				}
				for t > 0 {
					result = append(result, []int{-target, nums[j], nums[j]})
					break
				}
				m[nums[j]] = 0
			} else {
				f, s := nums[j], target-nums[j]
				total := m[f] * m[s]
				for total > 0 {
					result = append(result, []int{-target, f, s})
					break
				}
				m[f] = 0
				m[s] = 0
			}
		}
	}

	return result
}

// 724 ms	13.8 MB
func threeSumFirst(nums []int) [][]int {
	var result [][]int
	repeatMap := map[string]bool{}
	for i := 0; i < len(nums)-2; i++ {
		target := -nums[i]
		m := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			m[nums[j]] += 1
		}

		for j := i + 1; j < len(nums); j++ {
			if m[target-nums[j]] == 0 {
				continue
			}

			if nums[j]*2 == target {
				n, t := m[nums[j]], 0
				for n = n - 1; n > 0; n-- {
					t += n
				}
				for t > 0 {
					max, min := getMinMax(-target, nums[j], nums[j])
					k := fmt.Sprintf("%d%d", max, min)
					if !repeatMap[k] {
						result = append(result, []int{-target, nums[j], nums[j]})
						repeatMap[k] = true
					}
					break
				}
				m[nums[j]] = 0
			} else {
				f, s := nums[j], target-nums[j]
				total := m[f] * m[s]
				for total > 0 {
					max, min := getMinMax(-target, f, s)
					k := fmt.Sprintf("%d%d", max, min)
					if !repeatMap[k] {
						result = append(result, []int{-target, f, s})
						repeatMap[k] = true
					}
					break
				}
				m[f] = 0
				m[s] = 0
			}
		}
	}

	return result
}

func getMinMax(num1, num2, num3 int) (int, int) {
	return int(math.Max(math.Max(float64(num1), float64(num2)), float64(num3))),
		int(math.Min(math.Min(float64(num1), float64(num2)), float64(num3)))
}
