package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumTimeRequiredOvertime([]int{1, 2, 4, 7, 8}, 2))
	//fmt.Println(minimumTimeRequiredOvertime([]int{3, 2, 3}, 3))
	//fmt.Println(minimumTimeRequiredOvertime([]int{11, 2, 20, 18, 2, 1, 7, 11, 7, 10}, 9))
	//////[9899456,8291115,9477657,9288480,5146275,7697968,8573153,3582365,3758448,9881935,2420271,4542202]
	//fmt.Println(minimumTimeRequiredOvertime([]int{9899456, 8291115, 9477657, 9288480, 5146275, 7697968, 8573153, 3582365, 3758448, 9881935, 2420271, 4542202}, 9))
}

// 超时
func minimumTimeRequiredOvertime(jobs []int, k int) int {
	sort.Ints(jobs)
	totals := make([]int, k)
	n := len(jobs)
	ans := 0
	for i := 0; i < n && i < n/k+1; i++ {
		ans += jobs[n-i-1]
	}

	var dfs func(index int)
	dfs = func(index int) {
		if index < 0 {
			m := maxIndex(totals)
			if m < ans {
				ans = m
			}
			return
		}

		for i := 0; i < k; i++ {
			if totals[i]+jobs[index] < ans {
				totals[i] += jobs[index]
				dfs(index - 1)
				totals[i] -= jobs[index]
			}
		}
	}
	dfs(n - 1)
	return ans
}

func maxIndex(totals []int) int {
	m := totals[0]
	for i := 1; i < len(totals); i++ {
		if totals[i] > m {
			m = totals[i]
		}
	}
	return m
}
