package main

import "fmt"

func main() {
	//fmt.Println(numWays(3, 2), 4)
	//fmt.Println(numWays(2, 4), 2)
	//fmt.Println(numWays(4, 2), 8)
	fmt.Println(numWays(6, 13), 51)
}

// dp
func numWays(steps int, arrLen int) int {
	const mod = 1e9 + 7
	arrLen = min(steps, arrLen)
	dp := make([][]int, steps+1)
	for i := 0; i < steps+1; i++ {
		dp[i] = make([]int, arrLen)
	}
	dp[0][0] = 1

	for i := 1; i < steps+1; i++ {
		for j := 0; j < arrLen; j++ {
			count := dp[i-1][j]
			if j != 0 {
				count += dp[i-1][j-1]
			}
			if j != arrLen-1 {
				count += dp[i-1][j+1]
			}

			dp[i][j] = count % mod
		}
	}

	return dp[steps][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 剪枝
//func numWays(steps int, arrLen int) int {
//	dp := make([][]int, steps+1)
//	if arrLen > steps {
//		arrLen = steps
//	}
//	for i := 0; i < steps+1; i++ {
//		dp[i] = make([]int, arrLen)
//		for j := 0; j < arrLen && i != 0; j++ {
//			dp[i][j] = -1
//		}
//	}
//	dp[0][0] = 1
//
//	var search func(pos, leftSteps int) int
//	search = func(pos, leftSteps int) int {
//		if pos < 0 || pos >= arrLen {
//			return 0
//		} else if dp[leftSteps][pos] != -1 {
//			return dp[leftSteps][pos]
//		}
//
//		count := (search(pos, leftSteps-1) + search(pos-1, leftSteps-1) + search(pos+1, leftSteps-1)) % 1000000007
//		dp[leftSteps][pos] = count
//		return count
//	}
//
//	ans := search(0, steps)
//	return ans
//}
