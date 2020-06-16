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
	l4.Val = 2

	var l5 ListNode
	l5.Val = 1

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	result := isPalindrome(&l1)

	fmt.Println(result)
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return true
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil{
		//链表长度为奇数的话，取下一个节点为中间节点
		slow = slow.Next
	}

	var pre *ListNode
	for slow != nil{
		// 反转后半部分链表元素
		slow, pre, slow.Next = slow.Next, slow, pre
	}

	for pre != nil{
		if pre.Val != head.Val{
			return false
		}

		pre = pre.Next
		head = head.Next
	}

	return true
}