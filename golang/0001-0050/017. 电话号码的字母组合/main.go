package main

var dict = map[int32][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	stack := []string{""}

	for _, digit := range digits {
		newStack := make([]string, 0, len(stack)*4*len(digits))
		for _, l := range dict[digit] {
			for _, comb := range stack {
				newStack = append(newStack, comb+l)
			}
		}
		stack = newStack
	}

	return stack
}
