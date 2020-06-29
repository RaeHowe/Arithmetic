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
	l3.Val = 8

	var l4 ListNode
	l4.Val = 2

	var l5 ListNode
	l5.Val = 3

	var l6 ListNode
	l6.Val = 8

	l1.Next = &l2
	l2.Next = &l3

	l4.Next = &l5
	l5.Next = &l6

	result := mergeTwoLists(&l1, &l4)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Next.Val)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}else if l2 == nil{
		return l1
	}

	var dummyListNode = &ListNode{}
	var pre = dummyListNode

	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			pre.Next = l1
			pre = l1
			l1 = l1.Next
		}else {
			pre.Next = l2
			pre = l2
			l2 = l2.Next
		}
	}

	if l1 == nil{
		pre.Next = l2
	}else if l2 == nil{
		pre.Next = l1
	}

	return dummyListNode.Next
}