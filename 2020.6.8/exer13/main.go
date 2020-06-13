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
	l2.Val = 1

	var l3 ListNode
	l3.Val = 2

	l1.Next = &l2
	l2.Next = &l3

	result := deleteDuplicates(&l1)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
}

func deleteDuplicates(head *ListNode) *ListNode {
	var tmpMap = make(map[int]struct{})

	var dummyNode = &ListNode{}
	var preNode = dummyNode
	preNode.Next = head

	for head != nil{
		next := head.Next
		if _, ok := tmpMap[head.Val]; !ok{
			tmpMap[head.Val] = struct{}{}

			preNode = head
		}else {
			//重复元素
			preNode.Next = next
		}

		head = head.Next
	}

	return dummyNode.Next
}