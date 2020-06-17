package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
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

	result := oddEvenList(&l1)
	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Val)
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var odd = head
	var even = head.Next
	var evenHead = even

	for even != nil && even.Next != nil{
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}