package main

import (
	"sort"
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	n := len(words)
	wordsIndex := make(map[string]int, n)
	for i, word := range words {
		wordsIndex[word] = i
	}

	sort.Slice(words, func(i, j int) bool {
		wordI, wordJ := words[i], words[j]
		if len(wordI) == len(wordJ) {
			return wordsIndex[wordI] < wordsIndex[wordJ]
		}

		return len(wordI) < len(wordJ)
	})

	letterMap := make(map[int32]int)
	count := 0
	for _, l := range strings.ToLower(licensePlate) {
		if l >= 'a' && l <= 'z' {
			letterMap[l]++
			count++
		}
	}

	for _, word := range words {
		if len(word) < count {
			continue
		}

		wordLetterMap := make(map[int32]int)

		for _, l := range strings.ToLower(word) {
			if l >= 'a' && l <= 'z' {
				wordLetterMap[l]++
			}
		}

		ok := true
		for l, num := range letterMap {
			if num > wordLetterMap[l] {
				ok = false
				break
			}
		}

		if ok {
			return word
		}
	}

	return ""
}
