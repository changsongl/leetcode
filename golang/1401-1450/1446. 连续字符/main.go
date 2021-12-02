package main

func maxPower(s string) int {
	max, count := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}

		if count > max {
			max = count
		}
	}

	return max
}
