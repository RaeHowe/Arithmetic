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
	l4.Val = 3

	var l5 ListNode
	l5.Val = 2

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	result := removeDuplicateNodes(&l1)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
}

//双指针法
func removeDuplicateNodes2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow := head

	for slow != nil {
		fast := slow
		for fast.Next != nil {
			if slow.Val == fast.Next.Val {
				fast.Next = fast.Next.Next
			} else {
				fast = fast.Next
			}
		}
		slow = slow.Next
	}

	return head
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	var tmpMap = make(map[int]struct{})

	var tmpNode = &ListNode{}
	var pre = tmpNode

	for head != nil{
		if _, ok := tmpMap[head.Val]; !ok{
			//如果元素没有重复，就把该元素放置到map里面去，并且切换索引到下一个元素
			tmpMap[head.Val] = struct{}{}
			pre.Next = head //把原始的head链表给到pre节点后面
			pre = head //pre往后移动
		}else {
			pre.Next = head.Next //删除链表节点pre后面的节点
		}

		head = head.Next
	}

	return tmpNode.Next
}