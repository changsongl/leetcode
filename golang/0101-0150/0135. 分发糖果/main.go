package main

import "fmt"

func main() {
	fmt.Println(candy([]int{1, 2, 2, 8, 5, 7, 6, 9}))
	//fmt.Println(candy([]int{1, 0, 2}))
	// 1, 2, 1, 2
}

func candy(ratings []int) int {
	n := len(ratings)
	ans, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}

func candyTimeNSpaceN(ratings []int) (ans int) {
	n := len(ratings)
	left := make([]int, n)
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candyMySolution(ratings []int) int {
	ans := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		getRating(ratings, ans, i)
	}

	//fmt.Println(ans)
	sum := 0
	for i := 0; i < len(ans); i++ {
		sum += ans[i]
	}
	return sum
}

func getRating(ratings, ans []int, i int) int {
	if ans[i] != 0 {
		return ans[i]
	}

	num := 1
	if i != 0 && ratings[i] > ratings[i-1] {
		num = ans[i-1] + 1
	}

	if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
		num2 := getRating(ratings, ans, i+1) + 1
		if num2 > num {
			num = num2
		}
	}

	ans[i] = num
	return ans[i]
}
