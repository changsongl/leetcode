package main

import "math"

func poorPigs(buckets, minutesToDie, minutesToTest int) int {
	states := minutesToTest/minutesToDie + 1
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(states))))
}

func poorPigsDP(buckets, minutesToDie, minutesToTest int) int {
	if buckets == 1 {
		return 0
	}

	combinations := make([][]int, buckets+1)
	for i := range combinations {
		combinations[i] = make([]int, buckets+1)
	}
	combinations[0][0] = 1

	iterations := minutesToTest / minutesToDie
	f := make([][]int, buckets)
	for i := range f {
		f[i] = make([]int, iterations+1)
	}
	for i := 0; i < buckets; i++ {
		f[i][0] = 1
	}
	for j := 0; j <= iterations; j++ {
		f[0][j] = 1
	}

	for i := 1; i < buckets; i++ {
		combinations[i][0] = 1
		for j := 1; j < i; j++ {
			combinations[i][j] = combinations[i-1][j-1] + combinations[i-1][j]
		}
		combinations[i][i] = 1
		for j := 1; j <= iterations; j++ {
			for k := 0; k <= i; k++ {
				f[i][j] += f[k][j-1] * combinations[i][i-k]
			}
		}
		if f[i][iterations] >= buckets {
			return i
		}
	}
	return 0
}
