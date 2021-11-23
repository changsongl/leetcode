package main

func buddyStrings(s string, goal string) bool {
	n := len(s)

	if len(goal) != n {
		return false
	}

	if s == goal {
		dict := make(map[int32]bool)
		for _, c := range s {
			if dict[c] {
				return true
			}

			dict[c] = true
		}
		return false
	}

	dict := make([]int, 0, 2)
	for i := 0; i < n; i++ {
		if s[i] != goal[i] {
			dict = append(dict, i)
		}
	}

	if len(dict) != 2 {
		return false
	}

	return s[dict[0]] == goal[dict[1]] && s[dict[1]] == goal[dict[0]]
}
