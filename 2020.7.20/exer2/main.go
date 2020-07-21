package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 7

	var l2 ListNode
	l2.Val = 2

	var l3 ListNode
	l3.Val = 4

	var l4 ListNode
	l4.Val = 3

	var l5 ListNode
	l5.Val = 5

	var l6 ListNode
	l6.Val = 6

	var l7 ListNode
	l7.Val = 4

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4

	l5.Next = &l6
	l6.Next = &l7

	result := addTwoNumbers(&l1, &l5)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1NewHead = reverseListNode(l1)
	var l2NewHead = reverseListNode(l2)

	var dummyNode = &ListNode{}

	var pre = dummyNode

	var flag int
	for l1NewHead != nil || l2NewHead != nil{
		var num1, num2 int

		if l1NewHead != nil{
			num1 = l1NewHead.Val
			l1NewHead = l1NewHead.Next
		}

		if l2NewHead != nil{
			num2 = l2NewHead.Val
			l2NewHead = l2NewHead.Next
		}

		var sum = num1 + num2 + flag

		if sum >= 10{
			flag = 1
		}else {
			flag = 0
		}

		var tmpNode = &ListNode{
			Val: sum % 10,
		}

		pre.Next = tmpNode
		pre = tmpNode
	}

	if flag == 1 {
		var tmpNode = &ListNode{
			Val: 1,
		}
		pre.Next = tmpNode
	}

	return reverseListNode(dummyNode.Next)
}

func reverseListNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var tmpNode = reverseListNode(head.Next)
	head.Next.Next = head
	head.Next = nil

	return tmpNode
}
