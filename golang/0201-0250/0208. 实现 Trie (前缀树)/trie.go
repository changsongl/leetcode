package main

import "fmt"

func main() {
	prefix := "xxx"
	word := prefix + "haha"

	obj := Constructor()
	obj.Insert(word)
	fmt.Println(obj.Search(word))
	fmt.Println(obj.StartsWith(prefix))
	fmt.Println(obj.StartsWith("he"))
}

type Trie struct {
	let []*Trie
	end bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{let: make([]*Trie, 32), end: false}
}

/** Inserts a word into the trie. */
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

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this
	for _, let := range word {
		if t.let[let-'a'] == nil {
			return false
		}
		t = t.let[let-'a']
	}
	return t.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for _, let := range prefix {
		if t.let[let-'a'] == nil {
			return false
		}
		t = t.let[let-'a']
	}
	return true
}
