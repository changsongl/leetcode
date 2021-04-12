package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
}

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

//func largestNumber(nums []int) string {
//	sort.Slice(nums, func(i, j int) bool {
//		iStr, jStr := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
//		iFirst, jFirst := iStr+jStr, jStr+iStr
//		for i, let := range iFirst {
//			if let > int32(jFirst[i]) {
//				return true
//			} else if let < int32(jFirst[i]) {
//				return false
//			}
//		}
//		return true
//	})
//
//	var buffer bytes.Buffer
//
//	for _, num := range nums {
//		buffer.WriteString(strconv.Itoa(num))
//	}
//
//	ans := buffer.String()
//	if strings.Replace(ans, "0", "", -1) == "" {
//		return "0"
//	}
//
//	return ans
//}
