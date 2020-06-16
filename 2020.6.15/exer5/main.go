package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	var l1 ListNode
	l1.Val = 1

	var l2 ListNode
	l2.Val = 2

	var l3 ListNode
	l3.Val = 3

	var l4 ListNode
	l4.Val = 4

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4

	result := swapPairs(&l1)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var dummyNode = &ListNode{}
	var pre = dummyNode
	pre.Next = head
	
	for head != nil && head.Next != nil{
		pre.Next = head.Next
		head.Next = head.Next.Next
		pre.Next.Next = head

		pre = head
		head = head.Next
	}

	return dummyNode.Next
}