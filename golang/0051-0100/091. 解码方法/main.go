package main

import "fmt"

func main() {
	//fmt.Println(numDecodings("0"), 0)
	fmt.Println(numDecodings("12"), 2)
	fmt.Println(numDecodings("123"), 3)
	fmt.Println(numDecodings("0"), 0)
}

var m = map[string]bool{
	"1":  true,
	"2":  true,
	"3":  true,
	"4":  true,
	"5":  true,
	"6":  true,
	"7":  true,
	"8":  true,
	"9":  true,
	"10": true,
	"11": true,
	"12": true,
	"13": true,
	"14": true,
	"15": true,
	"16": true,
	"17": true,
	"18": true,
	"19": true,
	"20": true,
	"21": true,
	"22": true,
	"23": true,
	"24": true,
	"25": true,
	"26": true,
}

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	var dfs func(start int) int
	dfs = func(start int) int {
		if n-start == 0 {
			return 1
		} else if dp[start] != -1 {
			return dp[start]
		}

		count := 0
		if m[s[start:start+1]] {
			dp[start+1] = dfs(start + 1)
			count += dp[start+1]
		}

		if start < n-1 && m[s[start:start+2]] {
			dp[start+2] = dfs(start + 2)
			count += dp[start+2]
		}

		dp[start] = count
		return count
	}

	count := dfs(0)
	return count
}
