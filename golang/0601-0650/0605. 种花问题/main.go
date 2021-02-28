package main

// greedy algorithm
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n != 0; i += 2 {
		if flowerbed[i] == 0 {
			if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
				n--
			} else {
				i++
			}
		}
	}

	return n == 0
}
