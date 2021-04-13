package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(minNumber([]int{1, 0}))
}

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		i, j, is, js := nums[i], nums[j], 10, 10
		for ic, jc := i/10, j/10; ic != 0 || jc != 0; ic, jc = ic/10, jc/10 {
			if ic != 0 {
				is *= 10
			}

			if jc != 0 {
				js *= 10
			}
		}
		return i*js+j < j*is+i
	})

	var buffer bytes.Buffer
	for _, num := range nums {
		buffer.WriteString(strconv.Itoa(num))
	}
	return buffer.String()
}
