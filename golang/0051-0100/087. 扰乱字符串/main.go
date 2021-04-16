package main

import "fmt"

func main() {
	fmt.Println(isScramble("great", "rgeat"))
	fmt.Println(isScramble("abc", "bca"))
}

func isScramble(s1, s2 string) bool {
	n := len(s1)
	dp := make([][][]int8, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int8, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int8, n+1)
			for k := 0; k < n+1; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(i, j, l int) int8
	dfs = func(i1, j1, l int) (res int8) {
		d := &dp[i1][j1][l]
		if *d != -1 {
			return *d
		}
		defer func() { *d = res }()

		//fmt.Println(s1[i1:i1+l], s2[j1:j1+l], s1[i1:i1+l] == s2[j1:j1+l])
		if s1[i1:i1+l] == s2[j1:j1+l] {
			return 1
		}

		dict := [26]int{}
		for i := 0; i < n; i++ {
			dict[s1[i]-'a']++
			dict[s2[i]-'a']--
		}

		failed := false
		for i := 0; i < 26; i++ {
			if dict[i] != 0 {
				failed = true
				dict[i] = 0
			}
		}

		if failed {
			return 0
		}

		for i := 1; i < l; i++ {
			if dfs(i1, j1, i) == 1 && dfs(i1+i, j1+i, l-i) == 1 {
				return 1
			} else if dfs(i1, j1+l-i, i) == 1 && dfs(i1+i, j1, l-i) == 1 {
				return 1
			}
		}

		return 0
	}

	return dfs(0, 0, n) == 1
}

//func isScrambleBruteForce(s1 string, s2 string) bool {
//	if s1 == s2 {
//		return true
//	}
//
//	dict := make(map[byte]int, 26)
//	n := len(s1)
//	for i := 0; i < n; i++ {
//		dict[s1[i]]++
//		dict[s2[i]]--
//	}
//
//	for i := 0; i < 26; i++ {
//		if dict[byte('a'+i)] != 0 {
//			return false
//		}
//	}
//
//	for i := 1; i < n; i++ {
//		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
//			return true
//		} else if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
//			return true
//		}
//	}
//	return false
//}
