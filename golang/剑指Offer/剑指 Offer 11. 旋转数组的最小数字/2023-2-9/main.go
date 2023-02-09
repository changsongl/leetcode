package main

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1

	for left < right {
		mid := (left + right) / 2

		if numbers[right] > numbers[mid] {
			right = mid
		} else if numbers[right] < numbers[mid] {
			left = mid + 1
		} else {
			right--
		}
	}

	return numbers[left]
}
