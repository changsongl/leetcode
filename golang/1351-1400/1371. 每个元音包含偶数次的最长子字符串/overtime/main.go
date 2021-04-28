package main

import "fmt"

func main() {
	fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
}

func findTheLongestSubstring(s string) int {
	a, e, i, o, u := 0, 0, 0, 0, 0
	for _, l := range s {
		a, e, i, o, u = calAEIOU(a, e, i, o, u, byte(l), true)
	}

	var check func(l, r, a, e, i, o, u int) int
	check = func(l, r, a, e, i, o, u int) int {
		if a%2 == 0 && e%2 == 0 && i%2 == 0 && o%2 == 0 && u%2 == 0 {
			return r - l + 1
		} else if l > r {
			return 0
		}

		left, right := 0, 0

		start := l
		for start <= r {
			if s[start] == 'a' || s[start] == 'e' || s[start] == 'i' || s[start] == 'o' || s[start] == 'u' {
				ac, ec, ic, oc, uc := calAEIOU(a, e, i, o, u, s[start], false)
				left = check(start+1, r, ac, ec, ic, oc, uc)
				break
			}
			start++
		}

		end := r
		for l <= end {
			if s[end] == 'a' || s[end] == 'e' || s[end] == 'i' || s[end] == 'o' || s[end] == 'u' {
				ac, ec, ic, oc, uc := calAEIOU(a, e, i, o, u, s[end], false)
				right = check(l, end-1, ac, ec, ic, oc, uc)
				break
			}
			end--
		}

		if left > right {
			return left
		} else {
			return right
		}
	}

	return check(0, len(s)-1, a, e, i, o, u)
}

func calAEIOU(a, e, i, o, u int, l byte, add bool) (int, int, int, int, int) {
	switch l {
	case 'a':
		if add {
			a++
		} else {
			a--
		}
	case 'e':
		if add {
			e++
		} else {
			e--
		}
	case 'i':
		if add {
			i++
		} else {
			i--
		}
	case 'o':
		if add {
			o++
		} else {
			o--
		}
	case 'u':
		if add {
			u++
		} else {
			u--
		}
	}
	return a, e, i, o, u
}
