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
	l3.Val = 4

	var x1 ListNode
	x1.Val = 1

	var x2 ListNode
	x2.Val = 3

	var x3 ListNode
	x3.Val = 4

	l1.Next = &l2
	l2.Next = &l3

	x1.Next = &x2
	x2.Next = &x3

	result := mergeTwoLists(&l1, &x1)
	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Next.Val)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}else if l2 == nil{
		return l1
	}

	var tmpMidNode = &ListNode{}
	var pre = tmpMidNode

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

	return tmpMidNode.Next
}