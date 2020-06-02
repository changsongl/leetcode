package main

type ListNode struct {
	Val int
	Next *ListNode
}

//执行用时 : 8 ms , 在所有 Go 提交中击败了 93.74% 的用户
//内存消耗 : 4.7 MB , 在所有 Go 提交中击败了 93.10% 的用户
func addTwoNumbersBetterSpace(l1 *ListNode, l2 *ListNode) *ListNode {
	adder := 0
	result := l1
	var lastL1 *ListNode

	for {
		if l1 == nil {
			lastL1.Next = l2
			break
		}else if l2 == nil{
			break
		}

		l1.Val += l2.Val
		// 进位
		if adder == 1 {
			l1.Val++
			adder = 0
		}

		// 检查是否需要进位
		if l1.Val > 9 {
			l1.Val = l1.Val % 10
			adder = 1
		}

		lastL1 = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	if adder == 1 {
		for lastL1.Next != nil {
			if lastL1.Next.Val == 9 {
				lastL1.Next.Val = 0
				lastL1 = lastL1.Next
			}else{
				lastL1.Next.Val++
				adder = 0
				break
			}
		}

		if lastL1.Next == nil && adder == 1{
			lastL1.Next = &ListNode{1, nil}
		}
	}

	return result
}

//执行用时 : 28 ms , 在所有 Go 提交中击败了 5.89% 的用户
//内存消耗 : 4.7 MB , 在所有 Go 提交中击败了 93.10% 的用户
func addTwoNumbersBruteForce(l1 *ListNode, l2 *ListNode) *ListNode {
	adder := 0
	result := l1
	var lastL1, lastL2 *ListNode

	for {
		if l1 == nil && l2 == nil{
			break
		}else if l1 == nil && adder == 1{
			lastL1.Next = &ListNode{1, nil}
			l1 = lastL1.Next
			adder = 0
		}else if l2 == nil && adder == 1{
			lastL2.Next = &ListNode{1, nil}
			l2 = lastL2.Next
			adder = 0
		}else if l1 == nil {
			lastL1.Next = l2
			break
		}else if l2 == nil {
			break
		}

		l1.Val += l2.Val
		// 进位
		if adder == 1 {
			l1.Val++
			adder = 0
		}

		// 检查是否需要进位
		if l1.Val > 9 {
			l1.Val = l1.Val % 10
			adder = 1
		}

		lastL1 = l1
		lastL2 = l2
		l1 = l1.Next
		l2 = l2.Next
	}

	if adder == 1 {
		lastL1.Next = &ListNode{1, nil}
	}

	return result
}