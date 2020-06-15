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

	var l4 ListNode
	l4.Val = 1

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4

	result := isPalindrome(&l1)
	fmt.Println(result)
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return true
	}

	if head.Next.Next == nil{
		if head.Val == head.Next.Val{
			return true
		}else {
			return false
		}
	}

	var tmpArray []int

	for head != nil{
		tmpArray = append(tmpArray, head.Val)

		head = head.Next
	}

	var indexOfHead = 0
	var indexOfEnd = len(tmpArray) - 1
	var middleIndex = len(tmpArray) / 2


	var flag = true
	for i := 0; i < middleIndex; i++{
		if tmpArray[indexOfHead] == tmpArray[indexOfEnd]{
			indexOfHead++
			indexOfEnd--
		}else {
			flag = false
			break
		}
	}

	return flag
}