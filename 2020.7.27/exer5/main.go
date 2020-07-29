package main

type MyQueue struct {
	push []int
	pop []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.push = append(this.push, x)
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.pop) == 0{
		//如果弹出栈中的元素个数为0的话，就从压入栈取元素放置到这里
		for i := len(this.push) - 1; i >= 0; i--{
			this.pop = append(this.pop, this.push[i])
			this.push = this.push[:i]
		}
	}

	targetNumber := this.pop[len(this.pop) - 1]

	return targetNumber
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.pop) == 0 && len(this.push) == 0
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.pop) == 0{
		//如果弹出栈中的元素个数为0的话，就从压入栈取元素放置到这里
		for i := len(this.push) - 1; i >= 0; i--{
			this.pop = append(this.pop, this.push[i])
			this.push = this.push[:i]
		}
	}

	//放置完毕之后取弹出栈的第一个元素返回
	targetNumber := this.pop[len(this.pop) - 1]
	this.pop = this.pop[:len(this.pop) - 1]
	return targetNumber
}