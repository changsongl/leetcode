package main

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	dict map[int]*Node
	head *Node
	tail *Node
	cap  int
	size int
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next, tail.prev = tail, head
	return LRUCache{dict: make(map[int]*Node, capacity), cap: capacity, head: head, tail: tail}
}

func (this *LRUCache) destroy() {
	if this.cap < this.size && this.cap > 0 {
		d := this.tail.prev
		d.prev.next = this.tail
		this.tail.prev = d.prev
		delete(this.dict, d.key)
	}
}

func (this *LRUCache) update(node *Node, isNew bool) {
	if !isNew {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	this.head.next.prev = node
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
}

func (this *LRUCache) Get(key int) int {
	if n, exist := this.dict[key]; exist {
		this.update(n, false)
		return n.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, exist := this.dict[key]; exist {
		n.val = value
		this.update(n, false)
	} else {
		n := &Node{key: key, val: value}
		this.dict[key] = n
		this.size++
		this.update(n, true)
		this.destroy()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
