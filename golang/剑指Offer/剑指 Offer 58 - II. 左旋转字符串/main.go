package main

func reverseLeftWords1(s string, n int) string {
	bts := make([]byte, len(s))

	for i := range bts {
		bts[i] = s[(i+n)%len(s)]
	}

	return string(bts)
}

func main() {
	reverseLeftWords("abcde", 2)
}

func reverseLeftWords(s string, n int) string {
	l := len(s)

	bts := make([]byte, l)
	for i := range s {
		bts[i] = s[i]
	}

	reverseString(bts, 0, l-1)
	reverseString(bts, 0, l-n-1)
	reverseString(bts, l-n, l-1)

	return string(bts)
}

func reverseString(bts []byte, start, end int) {
	for i := start; i <= (start+end)/2; i++ {
		bts[i], bts[start+end-i] = bts[start+end-i], bts[i]
	}
}

// a b c d e 2 -> c d e a b
//  a b c d e
// e d c b a
// c d e b a
// c d e a b
