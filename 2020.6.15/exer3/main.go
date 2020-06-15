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
	l4.Val = 4

	var l5 ListNode
	l5.Val = 5

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	result := removeElements(&l1, 4)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}

	var dummyNode = &ListNode{}
	var pre = dummyNode
	pre.Next= head

	for head != nil{
		if head.Val == val{
			pre.Next = head.Next
		}else {
			pre = head
		}
		head = head.Next
	}

	return dummyNode.Next
}