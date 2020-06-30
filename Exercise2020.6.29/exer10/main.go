package main

import (
	"fmt"
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
	l3.Val = 1

	l1.Next = &l2
	l2.Next = &l3

	result := getDecimalValue(&l1)

	fmt.Println(result)

}

func getDecimalValue(head *ListNode) int {
	if head == nil{
		return 0
	}

	//链表反转
	newHead := reverseList(head)

	var result int

	var index = 0
	for newHead != nil{
		var tmpData = newHead.Val * pow(2, index)
		result += tmpData

		newHead = newHead.Next
		index++
	}

	return result
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var tmpListNode = reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tmpListNode
}

func pow(a, b int) int {
	var result = 1
	for i := 0; i < b; i++{
		result = result * a
	}

	return result
}