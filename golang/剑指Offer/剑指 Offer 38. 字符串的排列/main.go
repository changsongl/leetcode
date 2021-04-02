package main

import "fmt"

func main() {
	fmt.Println(permutation("abc"))
}

func permutation(s string) []string {
	chars := []byte(s)
	set := map[string]bool{}

	dfs(chars, 0, set)

	ans := make([]string, 0, len(set))
	for key := range set {
		ans = append(ans, key)
	}
	return ans
}

func dfs(chars []byte, index int, set map[string]bool) {
	if index == len(chars)-1 {
		set[string(chars)] = true
		return
	}

	for i := index; i < len(chars); i++ {
		swap(chars, index, i)
		dfs(chars, index+1, set)
		swap(chars, i, index)
	}
}

func swap(chars []byte, i1, i2 int) {
	chars[i1], chars[i2] = chars[i2], chars[i1]
}
