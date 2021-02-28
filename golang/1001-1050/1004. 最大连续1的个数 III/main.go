package main

import "fmt"

func main() {
	//fmt.Println(longestOnes([]int{0, 0, 0, 0}, 0))                                              // 0
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)) // 6
	//fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)) // 10
}

func longestOnes(A []int, K int) (ans int) {
	var left, lsum, rsum int
	for right, v := range A {
		rsum += 1 - v
		for lsum < rsum-K {
			fmt.Println(lsum, rsum, K)
			lsum += 1 - A[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestOnesOther(A []int, K int) int {
	n := len(A)         // 数组/字符串长度
	ans := 0            // 答案，最大饱满题目要求的长度
	left, right := 0, 0 // 双指针[left, right]
	zeros := 0          // 用于统计 子数组/子区间 是否有效，根据题目可能会改成求和/计数

	for right < n { // 循环没有到结尾时
		if A[right] == 0 { // 增加统计
			zeros += 1
		}

		for zeros > K { // 当不符合条件时候，移动左指针，调整有效统计
			if A[left] == 0 {
				zeros--
			}
			left++
		}

		// 检查结果是否需要更新
		if right-left+1 > ans {
			ans = right - left + 1
		}

		right++
	}

	return ans
}

// my solution
func longestOnesMySolution(A []int, K int) int {
	left, KCount, largest := 0, K, 0
	for i := 0; i < len(A); i++ {
		if A[i] == 1 {
			continue
		}

		if KCount == 0 {
			if i-left+1 > largest && i != left {
				largest = i - left
			}

			if K == 0 {
				left = i + 1
			} else {
				if A[left] == 0 {
					left++
				} else {
					for j := left + 1; j < i; j++ {
						if A[j] == 0 {
							left = j + 1
							break
						}
					}
				}
			}
		} else if A[i] == 0 {
			KCount--
		}
	}

	if len(A)-left > largest {
		largest = len(A) - left
	}

	return largest
}
