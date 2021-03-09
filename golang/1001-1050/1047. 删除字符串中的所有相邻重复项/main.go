package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(S string) string {
	if len(S) < 2 {
		return S
	}

	bytes := []byte(S)
	i := 0
	for i+1 < len(bytes) {
		if bytes[i] == bytes[i+1] {
			bytes = append(bytes[:i], bytes[i+2:]...)
			if i != 0 {
				i--
			}
		} else {
			i++
		}
	}

	return string(bytes)
}
