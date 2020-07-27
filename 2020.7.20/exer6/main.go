package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 3

	var l2 ListNode
	l2.Val = 2

	var l3 ListNode
	l3.Val = 0

	var l4 ListNode
	l4.Val = -4

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l2

	result := detectCycle(&l1)

	fmt.Println(result.Val)
}

func detectCycle(head *ListNode) *ListNode {
	var tmpMap = make(map[*ListNode]int)

	for head != nil{
		if _, ok := tmpMap[head]; ok {
			//如果存在
			return head
		}else {
			//如果不存在
			tmpMap[head]++
		}

		head = head.Next
	}

	return nil
}