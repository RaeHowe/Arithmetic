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

	var x1 ListNode
	x1.Val = 5

	var x2 ListNode
	x2.Val = 0

	var x3 ListNode
	x3.Val = 1

	x1.Next = &x2
	x2.Next = &x3
	x3.Next = &l3

	result := getIntersectionNode(&l1, &x1)
	fmt.Println(result.Val)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var tmpMap = make(map[*ListNode]int)

	for headA != nil || headB != nil{
		if headA != nil{
			if _, ok := tmpMap[headA]; !ok{
				tmpMap[headA] = headA.Val
			}else {
				return headA
			}

			headA = headA.Next
		}

		if headB != nil{
			if _, ok := tmpMap[headB]; !ok{
				tmpMap[headB] = headB.Val
			}else {
				return headB
			}

			headB = headB.Next
		}
	}

	return nil
}