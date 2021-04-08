package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{3, 1, 1, 1}))
}

// TODO：演示下
func singleNumber(nums []int) int {
	var twos, ones int
	for _, v := range nums {
		ones = ones ^ v&^twos
		twos = twos ^ v&ones
	}
	return ones
}

func singleNumberEasy(nums []int) int {
	dict := make([]int, 32)
	for _, num := range nums {
		i := 0
		for num != 0 {
			if num&1 == 1 {
				dict[i]++
			}
			i++
			num = num >> 1
		}
	}

	ans := 0
	for i := 31; i >= 0; i-- {
		ans = ans << 1
		if dict[i]%3 == 1 {
			ans += 1
		}
	}

	return ans
}

func singleNumberBruteForce(nums []int) int {
	dict := make(map[int]uint8)
	for _, num := range nums {
		dict[num]++
	}

	for _, num := range nums {
		if dict[num] == 1 {
			return num
		}
	}
	return 0
}
