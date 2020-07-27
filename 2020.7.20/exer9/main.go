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
	l2.Val = 1

	var l3 ListNode
	l3.Val = 6

	var l4 ListNode
	l4.Val = 5

	var l5 ListNode
	l5.Val = 9

	var l6 ListNode
	l6.Val = 2

	l1.Next = &l2
	l2.Next = &l3

	l4.Next = &l5
	l5.Next = &l6

	result := addTwoNumbers(&l1, &l4)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyNode = &ListNode{

	}

	var pre = dummyNode


	var flag int
	for l1 != nil || l2 != nil{
		var num1, num2 int

		if l1 != nil{
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil{
			num2 = l2.Val
			l2 = l2.Next
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

	if flag == 1{
		var tmpNode = &ListNode{
			Val: 1,
		}

		pre.Next = tmpNode
	}

	return dummyNode.Next
}

func reverseList(head *ListNode)  *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var tmpNode = reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return tmpNode
}