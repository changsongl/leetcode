package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

// log n solution
func maxArea(height []int) int {
	left, right, newArea := 0, len(height)-1, 0
	area := 0
	for left < right {
		if height[left] > height[right] {
			newArea = (right - left) * height[right]
			right--
		} else {
			newArea = (right - left) * height[left]
			left++
		}

		if area < newArea {
			area = newArea
		}
	}
	return area
}

// a little bit better filter 676 ms
func maxAreaALittleBetter(height []int) int {
	width, h, area, low, currentWidth := len(height), 0, 0, 0, 0
	for i := 0; i < width-1; i++ {
		for j := 0; j <= i; j++ {
			currentWidth = width - i - 1
			if height[j] == 0 || height[j+currentWidth] == 0 {
				continue
			}

			if height[j] < low {
				height[j] = 0
				continue
			} else if height[j+currentWidth] < low {
				height[j+currentWidth] = 0
				continue
			}

			left, right := height[j], height[j+currentWidth]
			if left > right {
				h = right
			} else {
				h = left
			}
			if area < h*currentWidth {
				low = h
				area = h * currentWidth
			}
		}
	}
	return area
}

// brute force 864 ms
func maxAreaBruteForce(height []int) int {
	width, h, area := len(height), 0, 0
	for i := 0; i < width-1; i++ {
		for j := 0; j <= i; j++ {
			currentWidth := width - i - 1
			left, right := height[j], height[j+currentWidth]

			if left > right {
				h = right
			} else {
				h = left
			}

			if area < h*currentWidth {
				area = h * currentWidth
			}
		}
	}

	return area
}
