package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) //3
	//fmt.Println(lengthOfLongestSubstring("abba"))     //2
}

func lengthOfLongestSubstring(s string) int {
	dict := make(map[int32]int)
	max := 0
	l := 0
	firstIndex := 1
	for i, c := range s {
		lastIndex := dict[c]
		if lastIndex == 0 || lastIndex < firstIndex {
			l++
		} else {
			l = i - lastIndex + 1
			firstIndex = lastIndex
		}
		dict[c] = i + 1

		max = Max(max, l)
	}
	return max
}

func Max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
