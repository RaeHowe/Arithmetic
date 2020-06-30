package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var array = []int{4, 9, 10, 2}

	result := convertArrayToList(array)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
	fmt.Println(result.Next.Next.Next.Val)
}

func convertArrayToList(tmpSlice []int) *ListNode {
	var dummyListNode = &ListNode{}
	var pre = dummyListNode

	for i := 0; i < len(tmpSlice); i++{
		var tmpNode = &ListNode{
			Val: tmpSlice[i],
		}

		pre.Next = tmpNode

		pre = tmpNode
	}

	return dummyListNode.Next
}