package main

import "fmt"

func main() {
	fmt.Println(numDistinct("rabbbit", "rabit"))
}

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	f := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		f[i] = make([]int, n+1)
		f[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				f[i][j] = f[i-1][j] + f[i-1][j-1]
			} else {
				f[i][j] = f[i-1][j]
			}
		}
	}
	return f[m][n]
}

type counter struct {
	pos   int
	count int
}

func numDistinctMySolution(s string, t string) int {
	count := 0
	var counterSlice []counter

	for i := 0; i < len(t); i++ {
		c := t[i]
		var counters []counter
		j := 0
		if len(counterSlice) == 0 {
			if i != 0 {
				break
			}
		} else {
			j = counterSlice[0].pos + 1
		}

		for ; j < len(s); j++ {
			match := s[j]
			if match == c {
				eleCount := 0
				if i == 0 {
					eleCount = 1
				} else {
					for _, counterEle := range counterSlice {
						if j > counterEle.pos {
							eleCount += counterEle.count
						}
					}
				}

				if eleCount > 0 {
					counters = append(counters, counter{
						pos:   j,
						count: eleCount,
					})
				}
			}

		}
		counterSlice = counters
	}
	for _, counter := range counterSlice {
		count += counter.count
	}
	return count
}
