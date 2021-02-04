package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("azxabbbavbn"))
}

// dp
func longestPalindrome(s string) string {
	ans := ""
	n := len(s)
	dp := make([][]int, n)

	// 初始化
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for l := 0; l < n; l++ {
		for i := 0; i < n-l; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] > 0 && len(ans) < l+1 {
				ans = s[i : j+1]
			}
		}
	}

	return ans
}
