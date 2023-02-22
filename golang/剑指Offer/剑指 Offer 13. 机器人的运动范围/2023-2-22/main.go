package main

func movingCount(m int, n int, k int) int {
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}

	var count int

	var f func(i, j int)
	f = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		} else if used[i][j] {
			return
		}

		used[i][j] = true

		if valid(i, j, k) {
			count++
			f(i+1, j)
			f(i-1, j)
			f(i, j+1)
			f(i, j-1)
		}
	}

	f(0, 0)

	return count
}

func valid(i, j, k int) bool {
	total := 0
	for i != 0 {
		total += i % 10
		i = i / 10
	}

	for j != 0 {
		total += j % 10
		j = j / 10
	}

	return k >= total
}
