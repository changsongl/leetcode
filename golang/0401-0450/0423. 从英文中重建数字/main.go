package main

import (
	"bytes"
	"fmt"
)

// zero      Z
// eight     g
// two       W
// six       X
// seven     S
// three     H
// four      R
// five      F
// one       O
// nine      I

func main() {
	fmt.Println(originalDigits("zerozero"))
}

type Dict struct {
	search    int32
	charCount map[int32]int
	num       int
}

var dicts = []Dict{
	{'z', map[int32]int{'z': 1, 'e': 1, 'r': 1, 'o': 1}, 0},
	{'g', map[int32]int{'i': 1, 'e': 1, 'g': 1, 'h': 1, 't': 1}, 8},
	{'w', map[int32]int{'w': 1, 'o': 1, 't': 1}, 2},
	{'x', map[int32]int{'i': 1, 's': 1, 'x': 1}, 6},
	{'s', map[int32]int{'s': 1, 'e': 2, 'v': 1, 'n': 1}, 7},
	{'h', map[int32]int{'t': 1, 'h': 1, 'r': 1, 'e': 2}, 3},
	{'r', map[int32]int{'f': 1, 'o': 1, 'u': 1, 'r': 1}, 4},
	{'f', map[int32]int{'f': 1, 'i': 1, 'v': 1, 'e': 1}, 5},
	{'o', map[int32]int{'o': 1, 'n': 1, 'e': 1}, 1},
	{'i', map[int32]int{'n': 2, 'i': 1, 'e': 1}, 9},
}

func originalDigits(s string) string {
	charsCountMap := make(map[int32]int, 26)
	for _, c := range s {
		charsCountMap[c]++
	}

	charSlice := make([]int, 10)
	for _, dict := range dicts {
		num := charsCountMap[dict.search]
		if num > 0 {
			charSlice[dict.num] = num
			for c, cNum := range dict.charCount {
				charsCountMap[c] -= cNum * num
			}
		}
	}

	var ans bytes.Buffer

	for c, num := range charSlice {
		if num > 0 {
			for i := 0; i < num; i++ {
				ans.WriteByte(byte('0' + c))
			}
		}
	}

	return ans.String()
}
