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

	result := getKthFromEnd(&l1, 2)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	//1 2 3 4 5
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head

	for k > 0{
		k--
		fast = fast.Next

		if fast == nil{
			return slow
		}
	}

	for fast != nil{
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

//把链表元素放置到数组里面
func getKthFromEnd3(head *ListNode, k int) *ListNode {
	var res []*ListNode
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	l := len(res)
	if l >= k {
		return res[l-k]
	}
	return nil
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	var cur = head
	newHead := reverseListNode(cur)

	var index = 1
	var result int
	for newHead != nil{
		if index == k{
			result = newHead.Val
		}

		newHead = newHead.Next
		index++
	}

	for head != nil{
		if head.Val == result{
			return head
		}

		head = head.Next
	}

	return nil
}

func reverseListNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var tmpNode = reverseListNode(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tmpNode
}