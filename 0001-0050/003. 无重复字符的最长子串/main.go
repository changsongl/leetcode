package main

import (
	"fmt"
)

func main(){
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

//执行用时 : 8 ms , 在所有 Go 提交中击败了 66.15% 的用户
//内存消耗 : 3.2 MB , 在所有 Go 提交中击败了 32.26% 的用户
func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen < 2 {
		return sLen
	}

	filters := make(map[int32]int)
	start, current, max := 0, 0, 0

	for i, c := range s {
		if i > 0 {
			current++
			lastI, exist := filters[c]
			if exist {
				if lastI >= start{
					if current - start > max {
						max = current - start
					}
					//做操作
					start = lastI + 1
				}
			}
		}
		filters[c] = i
	}

	if current - start + 1 > max {
		return current - start + 1
	}
	return max
}