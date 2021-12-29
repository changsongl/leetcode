package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findAllConcatenatedWordsInADict([]string{"", "cat", "catcat"}))
}

func findAllConcatenatedWordsInADict(words []string) []string {
	var ans []string

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	pTree := Constructor()

	for _, word := range words {
		if word == "" {
			continue
		}

		finished := pTree.dfs(word)
		if finished {
			ans = append(ans, word)
		} else {
			pTree.Insert(word)
		}
	}

	return ans
}

func (root *Trie) dfs(word string) bool {
	if word == "" {
		return true
	}
	node := root
	for i, ch := range word {
		node = node.let[ch-'a']
		if node == nil {
			return false
		}
		if node.end && root.dfs(word[i+1:]) {
			return true
		}
	}
	return false
}

type Trie struct {
	let []*Trie
	end bool
}

func Constructor() Trie {
	return Trie{let: make([]*Trie, 32), end: false}
}

func (this *Trie) Insert(word string) {
	t := this
	for _, let := range word {
		if t.let[let-'a'] == nil {
			newT := Constructor()
			t.let[let-'a'] = &newT
		}
		t = t.let[let-'a']
	}
	t.end = true
}
