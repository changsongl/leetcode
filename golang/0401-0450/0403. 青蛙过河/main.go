package main

func main() {
	//fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	//fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
	//fmt.Println(canCross([]int{0, 2}))
}

func canCross(stones []int) bool {
	n := len(stones)
	stoneMap := make(map[int]int, n)
	dp := make(map[int]map[int]bool, n)
	for i, stone := range stones {
		stoneMap[stone] = i
		dp[i] = make(map[int]bool)
	}

	var check func(index, step int) bool
	check = func(index, step int) bool {
		if index == n-1 {
			return true
		} else if dp[index][step] {
			return false
		}

		dp[index][step] = true
		nextStep := []int{step - 1, step, step + 1}
		for _, next := range nextStep {
			stone := stoneMap[stones[index]+next]
			if next > 0 && stone != 0 && check(stone, next) {
				return true
			}
		}
		return false
	}

	return check(0, 0)
}
