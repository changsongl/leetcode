package main

import "strconv"

func main() {

}

// 自己想到的，因为是回文，因此正反都一样。所以翻转之后相减应为0，
// 因为是回文，所以正常翻转后不会溢出。溢出的话也就不是回文了。
// 所以往证书回转上套，解决这个问题。
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xRev, xCopy := 0, x
	for xCopy != 0 {
		xRev = xRev*10 + xCopy%10
		xCopy /= 10
	}
	return xRev-x == 0
}

// 粗暴解法
func isPalindromeEasy(x int) bool {
	xStr := strconv.Itoa(x)
	var end, l = len(xStr) - 1, len(xStr) / 2
	for i := 0; i < l; i++ {
		if xStr[i] != xStr[end] {
			return false
		}
		end--
	}

	return true
}
