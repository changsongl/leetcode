package main

type pair struct {
	pos int
	ch  byte
}

func firstUniqChar(s string) byte {
	n := len(s)
	var queue []pair
	pos := [26]int{}
	for i := range pos {
		pos[i] = n
	}

	for i := range s {
		c := s[i]

		if pos[c-'a'] == n {
			queue = append(queue, pair{pos: i, ch: c})
			pos[c-'a'] = i
		} else {
			pos[c-'a'] = n + 1
			for len(queue) > 0 && pos[queue[0].ch-'a'] == n+1 {
				queue = queue[1:]
			}
		}
	}

	if len(queue) == 0 {
		return ' '
	}

	return queue[0].ch
}

func main() {
	firstUniqChar("leetcode")
}
