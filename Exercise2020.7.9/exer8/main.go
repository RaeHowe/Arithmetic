package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 4

	var l2 ListNode
	l2.Val = 5

	var l3 ListNode
	l3.Val = 1

	var l4 ListNode
	l4.Val = 9

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4

	result := deleteNode(&l1, 5)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}

	var dummyNode = &ListNode{
		Next: head,
	}

	var pre = dummyNode

	for head != nil{
		if head.Val == val{
			pre.Next = pre.Next.Next
		}

		pre = head
		head = head.Next
	}

	return dummyNode.Next
}
