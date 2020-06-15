package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
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
}

func loopList(head *ListNode) bool{
	if head == nil || head.Next == nil{
		return false
	}

	var slow = head
	var fast = head.Next

	for fast != nil && fast.Next != nil && slow != nil{
		if fast == slow{
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}