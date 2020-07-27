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
	l5.Val = 4

	var l6 ListNode
	l6.Val = 4

	var l7 ListNode
	l7.Val = 5

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6
	l6.Next = &l7

	result := deleteDuplicates2(&l1)

	fmt.Println(result.Val)
	fmt.Println(result.Next.Val)
	fmt.Println(result.Next.Next.Val)
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var dummyNode = &ListNode{
		Next: head,
	}

	var pre = dummyNode
	var behind = dummyNode.Next

	var flag = false

	for behind.Next != nil{
		if behind.Val == behind.Next.Val{
			flag = true
			behind = behind.Next
		}else {
			if flag {
				//目的是为了删除最后一个重复的那个元素
				behind = behind.Next
				pre.Next = behind
				flag = false //重置标识位
			}else {
				pre = pre.Next
				behind = behind.Next
			}
		}
	}

	if flag {
		pre.Next = behind.Next
	}

	return dummyNode.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	var tmpMap = make(map[int]int)

	var dummyNode = &ListNode{
		Next: head,
	}

	var pre = dummyNode

	for head != nil{
		if _, ok := tmpMap[head.Val]; ok {
			//存在的话，就删除该元素
			tmpMap[head.Val]++
			pre.Next = pre.Next.Next //删除下一个元素
		}else {
			//不存在的话就放置到map中
			tmpMap[head.Val]++
			pre = head
		}

		head = head.Next
	}

	var lastNode =  dummyNode.Next
	for lastNode != nil{
		if v, ok := tmpMap[lastNode.Val]; ok {
			if v > 1{
				lastNode.Val = lastNode.Next.Val
				lastNode.Next = lastNode.Next.Next
				continue
			}
		}

		lastNode = lastNode.Next
	}

	return dummyNode.Next
}
