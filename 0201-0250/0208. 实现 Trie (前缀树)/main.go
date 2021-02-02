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
	root *Node
}

type Node struct {
	val      rune
	term     bool
	children map[rune]*Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &Node{
			val:      '*',
			children: map[rune]*Node{},
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if this.Search(word) == true {
		return
	}

	n := this.root
	for len(word) > 0 {
		r := rune(word[0])
		word = word[1:]

		c, exist := n.children[r]
		if exist {
			n = c
		} else {
			newNode := &Node{
				val:      r,
				children: map[rune]*Node{},
			}
			n.children[r] = newNode
			n = newNode
		}

		if len(word) == 0 {
			n.term = true
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	n := this.root

	for len(word) > 0 {
		r := rune(word[0])
		word = word[1:]
		c, exist := n.children[r]
		if !exist {
			return false
		}

		if len(word) == 0 {
			return c.term
		}

		n = c
	}

	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	n := this.root

	for len(prefix) > 0 {
		r := rune(prefix[0])
		prefix = prefix[1:]
		c, exist := n.children[r]
		if !exist {
			return false
		}

		n = c
	}

	return true
}
