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
	l2.Val = 4

	var l3 ListNode
	l3.Val = 3

	var l4 ListNode
	l4.Val = 2

	var l5 ListNode
	l5.Val = 5

	var l6 ListNode
	l6.Val = 2

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6

	result := partition(&l1, 3)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Next.Val)
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil{
		return head
	}

	var dummyGreaterListNode = &ListNode{}

	var preGreater = dummyGreaterListNode

	var dummyLessListNode = &ListNode{}

	var preLess = dummyLessListNode

	for head != nil{
		if head.Val < x{
			preLess.Next = head
			preLess = head
		}else {
			preGreater.Next = head
			preGreater = head
		}

		head = head.Next
	}

	preGreater.Next = nil
	preLess.Next = dummyGreaterListNode.Next

	return dummyLessListNode.Next
}
