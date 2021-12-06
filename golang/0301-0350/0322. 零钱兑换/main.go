package main

import (
	"fmt"
)

func main() {
	fmt.Println(coinChange([]int{13, 16, 9, 1}, 27))
}

func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i := 0; i < max; i++ {
		dp[i] = max
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin && dp[i-coin]+1 <= dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

//func coinChange(coins []int, amount int) int {
//	if amount == 0 {
//		return 0
//	}
//
//	dp := make(map[int]int)
//	uncheck := coins
//	round := 1
//
//	for len(uncheck) != 0 {
//		newUncheck := make([]int, 0)
//		for _, check := range uncheck {
//			if dp[check] != 0 {
//				continue
//			} else if check > amount {
//				continue
//			} else if check == amount {
//				return round
//			} else {
//				dp[check] = round
//				for _, coin := range coins {
//					newUncheck = append(newUncheck, check+coin)
//				}
//			}
//		}
//
//		round++
//		uncheck = newUncheck
//	}
//
//	return -1
//}
