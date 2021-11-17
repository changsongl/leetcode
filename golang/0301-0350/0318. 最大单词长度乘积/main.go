package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}), 16)
	fmt.Println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}), 4)
	fmt.Println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}), 0)
}

func maxProduct(words []string) (ans int) {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}

	for i, x := range masks {
		for j, y := range masks[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return
}

func maxProductStraight(words []string) int {
	max := 0
	n := len(words)
	charCounters := make([][]int32, n)

	for i := range charCounters {
		charCounters[i] = make([]int32, 26)
	}

	for i, word := range words {
		for _, c := range word {
			charCounters[i][c-'a']++
		}
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {

			hasSameCh := false
			for ci := 0; ci < 26; ci++ {
				if charCounters[i][ci] != 0 && charCounters[j][ci] != 0 {
					hasSameCh = true
					break
				}
			}

			if hasSameCh {
				continue
			}

			prod := len(words[i]) * len(words[j])
			if prod > max {
				max = prod
			}
		}
	}

	return max
}
