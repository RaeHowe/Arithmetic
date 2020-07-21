package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 3

	var l2 ListNode
	l2.Val = 2

	var l3 ListNode
	l3.Val = 0

	var l4 ListNode
	l4.Val = -4

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4

	result := reversePrint(&l1)

	fmt.Println(result)
}

func reversePrint(head *ListNode) []int {
	var newHead = reverse(head)
	var resultArray []int

	for newHead != nil{
		resultArray = append(resultArray, newHead.Val)

		newHead = newHead.Next
	}

	return resultArray
}

func reverse(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}

	var tmpNode = reverse(head.Next)
	head.Next.Next = head
	head.Next = nil

	return tmpNode
}
