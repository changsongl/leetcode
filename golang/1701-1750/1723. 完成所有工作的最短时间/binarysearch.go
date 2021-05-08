package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(minimumTimeRequiredNormal([]int{1, 2, 4, 7, 8}, 2))
	//fmt.Println(minimumTimeRequiredNormal([]int{3, 2, 3}, 3))
	//fmt.Println(minimumTimeRequiredNormal([]int{11, 2, 20, 18, 2, 1, 7, 11, 7, 10}, 9))
	//[9899456,8291115,9477657,9288480,5146275,7697968,8573153,3582365,3758448,9881935,2420271,4542202]
	fmt.Println(minimumTimeRequiredNormal([]int{250, 250, 256, 251, 254, 254, 251, 255, 250, 252, 254, 255}, 10))
	fmt.Println(minimumTimeRequiredNormal([]int{9899456, 8291115, 9477657, 9288480, 5146275, 7697968, 8573153, 3582365, 3758448, 9881935, 2420271, 4542202}, 9))
}

// 超时
func minimumTimeRequiredNormal(jobs []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
	n := len(jobs)
	left, right := jobs[0], 0
	for _, job := range jobs {
		right += job
	}

	return left + sort.Search(right-left, func(i int) bool {
		limit := left + i
		workers := make([]int, k)
		var backtrace func(index int) bool
		backtrace = func(index int) bool {
			if index == n {
				return true
			}

			curJob := jobs[index]
			for i := 0; i < k; i++ {
				if workers[i]+curJob <= limit {
					workers[i] += curJob
					if backtrace(index + 1) {
						return true
					}
					workers[i] -= curJob
				}

				if workers[i] == 0 || workers[i]+curJob == limit {
					return false
				}
			}

			return false
		}

		return backtrace(0)
	})
}
