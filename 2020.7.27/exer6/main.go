package main

type MyQueue struct {
	list []int
}

func Constructor() MyQueue {
	return MyQueue{list: make([]int, 0)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.list = append(this.list, x)
}

