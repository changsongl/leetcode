package main

import "math"

func findShortestSubArray(nums []int) int {
	max, shortest := 0, math.MaxInt64
	type detail struct {
		count int
		start int
		end   int
	}

	m := make(map[int]detail, len(nums)/4)
	for i, num := range nums {
		d, ok := m[num]
		if ok {
			d.count++
			d.end = i
		} else {
			d = detail{
				count: 1,
				start: i,
				end:   i,
			}
		}
		m[num] = d

		if d.count > max {
			max = d.count
		}
	}

	for _, d := range m {
		if d.count == max && d.end-d.start+1 < shortest {
			shortest = d.end - d.start + 1
		}
	}

	return shortest
}
