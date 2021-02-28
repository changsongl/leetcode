package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}

var res []string

func generateParenthesis(n int) []string {
	res = nil

	generate("", 0, 0, n)

	return res
}

func generate(ans string, countL, countR, total int) {
	if countL > total || countR > total {
		return
	} else if countL == total && countR == total {
		res = append(res, ans)
	}

	if countL >= countR {
		generate(ans+"(", countL+1, countR, total)
		generate(ans+")", countL, countR+1, total)
	}
}
