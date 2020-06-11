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
	l4.Val = 4

	var l5 ListNode
	l5.Val = 5

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	result := middleNode(&l1)

	fmt.Println(result.Val)
}

func middleNode(head *ListNode) *ListNode {
	var tmpArray []*ListNode

	for head != nil{
		tmpArray = append(tmpArray, head)
		head = head.Next
	}

	var index = len(tmpArray) / 2
	return tmpArray[index]
}
