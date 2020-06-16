package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 9

	var l2 ListNode
	l2.Val = 1

	var l3 ListNode
	l3.Val = 3

	var l4 ListNode
	l4.Val = 5

	var l5 ListNode
	l5.Val = 2

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	result := sortList(&l1)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Val)
}

func sortList(head *ListNode) *ListNode {
	quickSort(head, nil)
	return head
}

func quickSort(head, end *ListNode)  {
	if head == end || head.Next == end{
		return
	}

	pivot := head.Val
	slow, fast := head, head.Next
	for fast != end{
		if fast.Val <= pivot{
			slow = slow.Next
			swap(slow, fast)
		}

		fast = fast.Next
	}

	swap(head, slow)
	quickSort(head, slow)
	quickSort(slow.Next, end)
}

func swap(a, b *ListNode)  {
	a.Val, b.Val = b.Val, a.Val
}