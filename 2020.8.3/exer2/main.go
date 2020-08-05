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
	l4.Val = -3

	var l5 ListNode
	l5.Val = 1

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	result := removeZeroSumSublists(&l1)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	var dummyNode = &ListNode{
		Next: head,
	}

	var pre = dummyNode

	for head != nil{
		var sum int

		for head != nil{
			sum += head.Val

			if sum == 0{
				break
			}

			head = head.Next
		}

		if sum == 0{
			head = head.Next
			pre.Next = head
		}else {
			pre = pre.Next
			head = pre.Next
		}
	}

	return dummyNode.Next
}