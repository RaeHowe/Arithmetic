package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 4

	var l2 ListNode
	l2.Val = 1

	var l3 ListNode
	l3.Val = 8

	var l4 ListNode
	l4.Val = 4

	var l5 ListNode
	l5.Val = 5

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	var l6 ListNode
	l6.Val = 5

	var l7 ListNode
	l7.Val = 0

	var l8 ListNode
	l8.Val = 1

	l6.Next = &l7
	l7.Next = &l8
	l8.Next = &l3

	result := getIntersectionNode(&l1, &l6)

	fmt.Println(result.Val)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var tmpMap = make(map[*ListNode]int)

	for headA != nil{
		tmpMap[headA]++
		headA = headA.Next
	}

	for headB != nil{
		if _, ok := tmpMap[headB]; ok{
			//如果存在的话
			return headB
		}else {
			headB = headB.Next
		}
	}

	return nil
}