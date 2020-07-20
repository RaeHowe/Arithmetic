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
	l2.Val = 8

	//var l3 ListNode
	//l3.Val = 3

	var l4 ListNode
	l4.Val = 0

	//var l5 ListNode
	//l5.Val = 6

	//var l6 ListNode
	//l6.Val = 4

	l1.Next = &l2
	//l2.Next = &l3

	//l4.Next = &l5
	//l5.Next = &l6

	result := addTwoNumbers(&l1, &l4)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return reverseListNode(l2)
	}

	if l2 == nil{
		return reverseListNode(l1)
	}

	l1NewHead := reverseListNode(l1)

	l2NewHead := reverseListNode(l2)

	var flag int

	var dummyNode = &ListNode{}
	var pre = dummyNode

	for l1NewHead != nil || l2NewHead != nil{
		var t1, t2 int
		if l1NewHead != nil{
			t1 = l1NewHead.Val
			l1NewHead = l1NewHead.Next
		}

		if l2NewHead != nil{
			t2 = l2NewHead.Val
			l2NewHead = l2NewHead.Next
		}

		var sum = t1 + t2 + flag

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

	var tmpListNode = reverseListNode(head.Next)
	head.Next.Next = head
	head.Next = nil

	return tmpListNode
}