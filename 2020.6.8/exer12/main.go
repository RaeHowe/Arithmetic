package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = -3

	var l2 ListNode
	l2.Val = 5

	var l3 ListNode
	l3.Val = -99

	//var l4 ListNode
	//l4.Val = 4
	//
	//var l5 ListNode
	//l5.Val = 5
	//
	//var l6 ListNode
	//l6.Val = 6

	l1.Next = &l2
	l2.Next = &l3


	result := deleteNode(&l1, -99)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	//fmt.Println(result.Next.Next.Val)
	//fmt.Println(result.Next.Next.Next.Val)
	//fmt.Println(result.Next.Next.Next.Next.Val)
}

func deleteNode(head *ListNode, val int) *ListNode {
	var dummyNode = &ListNode{
		Next: head,
	}

	var preNode = dummyNode

	for head != nil{
		next := head.Next
		if head.Val == val{
			preNode.Next = next
			break
		}

		preNode = head
		head = next
	}

	return dummyNode.Next
}