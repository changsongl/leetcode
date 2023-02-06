package main

import (
	"bytes"
)

func replaceSpace(s string) string {
	b := bytes.NewBufferString("")
	b.Grow(len(s))

	for _, c := range s {
		if c == ' ' {
			b.WriteString("%20")
		} else {
			b.WriteRune(c)
		}
	}

	return b.String()
}
