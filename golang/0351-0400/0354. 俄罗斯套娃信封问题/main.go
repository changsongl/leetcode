package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxEnvelopes([][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}}))
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] ||
			(envelopes[i][0] == envelopes[j][0] && envelopes[i][1] < envelopes[j][1])
	})

	l, ans := len(envelopes), 0
	dp := make([]int, len(envelopes))

	for i := 0; i < l; i++ {
		max := 1
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] && envelopes[i][0] != envelopes[j][0] && max < dp[j]+1 {
				max = dp[j] + 1
			}
		}

		dp[i] = max
		if max > ans {
			ans = max
		}
	}

	return ans
}
