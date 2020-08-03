package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 1

	var l2 ListNode
	l2.Val = 2

	var l3 ListNode
	l3.Val = 3

	var l4 ListNode
	l4.Val = 3

	var l5 ListNode
	l5.Val = 2

	var l6 ListNode
	l6.Val = 1

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6

	result := removeDuplicateNodes(&l1)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	var dummyNode = &ListNode{
		Next: head,
	}

	var pre = dummyNode

	var tmpMap = make(map[int]int)

	for head != nil{
		if _, ok := tmpMap[head.Val]; ok {
			//存在的话
			pre.Next = pre.Next.Next
		}else {
			//不存在的话
			tmpMap[head.Val]++

			pre = head
		}

		head = head.Next
	}

	return dummyNode.Next
}