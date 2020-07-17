package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 2

	var l2 ListNode
	l2.Val = 4

	var l3 ListNode
	l3.Val = 3

	var l4 ListNode
	l4.Val = 5

	var l5 ListNode
	l5.Val = 6

	var l6 ListNode
	l6.Val = 4

	l1.Next = &l2
	l2.Next = &l3

	l4.Next = &l5
	l5.Next = &l6

	result := addTwoNumbers(&l1, &l2)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

}