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
	l4.Val = 4

	var l5 ListNode
	l5.Val = 5

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	result := rotateRight(&l1, 2)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Next.Val)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0{
		return head
	}

	var count = 1
	var tmpNode = head

	for tmpNode.Next != nil{
		//统计链表节点总数量
		count++
		tmpNode = tmpNode.Next
	}

	//计算出移动次数(防止大于count的值)
	k = k % count
	//形成环形链表
	tmpNode.Next = head

	for i := 1; i <= count - k; i++{
		tmpNode = tmpNode.Next
	}

	head = tmpNode.Next
	tmpNode.Next = nil

	return head
}