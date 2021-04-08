package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}

func firstUniqCharBrute(s string) byte {
	m := make(map[byte]bool)
	chars := []byte(s)
	for _, b := range chars {
		ok, exists := m[b]
		if !exists {
			m[b] = false
		} else {
			if !ok {
				m[b] = true
			}
		}
	}

	for _, b := range chars {
		ok, _ := m[b]
		if !ok {
			return b
		}
	}

	return ' '
}

//func firstUniqChar(s string) byte {
//	dict := make([]*bool, 32)
//	for _, b := range s {
//		if dict[b-'a'] == nil {
//			cInd := false
//			dict[b-'a'] = &cInd
//		} else if *dict[b-'a'] == false {
//			cInd := true
//			dict[b-'a'] = &cInd
//		}
//	}
//
//	for _, b := range s {
//		if *dict[b-'a'] == false {
//			return byte(b)
//		}
//	}
//
//	return ' '
//}

func firstUniqChar(s string) byte {
	byteMap := [26]int{}
	for _, val := range s {
		byteMap[val-'a']++
	}
	for index, val := range s {
		if byteMap[val-'a'] == 1 {
			return s[index]
		}
	}
	return ' '
}
