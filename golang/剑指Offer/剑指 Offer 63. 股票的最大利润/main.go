package main

import "math"

func maxProfit(prices []int) int {
	minPrice, maxProf := math.MaxInt, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if maxProf < prices[i]-minPrice {
			maxProf = prices[i] - minPrice
		}
	}

	return maxProf
}

// 0    -6     4                 -2            3           -2
//       0 - 7 + 1    (-6 -1 + 5) = 2    4 - 5 + 3     2 - 3 + 6    5 - 6 + 4
func maxProfit2(prices []int) int {
	var max int
	dp := make([]int, len(prices))

	for i := 1; i < len(prices); i++ {
		dp[i] = prices[i] - prices[i-1]
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func maxProfit1(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}

	var max int
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			profit := prices[j] - prices[i]

			if profit > max {
				max = profit
			}
		}
	}

	return max
}
