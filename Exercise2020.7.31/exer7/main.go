package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 1

	var l2 ListNode
	l2.Val = 0

	var l3 ListNode
	l3.Val = 0

	var l4 ListNode
	l4.Val = 1

	var l5 ListNode
	l5.Val = 1

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	result := getDecimalValue(&l1)

	fmt.Println(result)
}

func getDecimalValue(head *ListNode) int {
	newHead := reverseList(head)

	var result int

	var flag = 0
	for newHead != nil{
		var tmpData = math.Pow(2, float64(flag))
		result += newHead.Val * int(tmpData)

		newHead = newHead.Next
		flag++
	}

	return result
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var tmpNode = reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return tmpNode
}