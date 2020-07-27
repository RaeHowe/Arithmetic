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

	result := removeNthFromEnd(&l1, 2)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//反转链表
	newHead := reverseList(head)

	var dummyNode = &ListNode{
		Next: newHead,
	}

	var pre = dummyNode

	var index = 1

	for newHead != nil{
		if index == n{
			pre.Next = pre.Next.Next
			break
		}else {
			pre = newHead
			index++
		}

		newHead = newHead.Next
	}

	return reverseList(dummyNode.Next)
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