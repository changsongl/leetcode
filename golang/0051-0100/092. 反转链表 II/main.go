package main

import "fmt"

func main() {
	h := reverseBetween(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}, 1, 4)
	printNodes(h)
}

func printNodes(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetweenMine(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	var front, back, reverseNode *ListNode
	node := head
	count := 1
	for node != nil {
		if count == left-1 {
			front = node
		} else if count == left {
			reverseNode = node
		} else if right == count {
			back = node.Next
			node.Next = nil
		}
		node = node.Next
		count++
	}

	//printNodes(front)
	//fmt.Println("--------")
	//printNodes(reverseNode)
	//fmt.Println("--------")
	//printNodes(back)
	//return nil

	rHead, rTail := reverse(reverseNode)
	if rTail.Next != nil {
		rTail.Next = back
	}

	if front == nil {
		head = rHead
	} else {
		front.Next = rHead
	}

	return head
}

func reverse(node *ListNode) (head *ListNode, tail *ListNode) {
	if node.Next == nil {
		return node, node
	}

	rHead, tail := reverse(node.Next)
	tail.Next = node
	return rHead, node
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	var prev *ListNode
	for m > 1 {
		prev = cur
		cur = cur.Next
		m--
		n--
	}
	con1, con2 := prev, cur
	for n > 0 {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		n--
	}
	if con1 == nil {
		con2.Next = cur
		return prev
	}
	con1.Next = prev
	con2.Next = cur
	return head
}
