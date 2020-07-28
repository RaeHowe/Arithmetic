package main

type CQueue struct {
	p []int
	q []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int)  {
	this.p = append(this.p, value) //往数组里面添加元素
}

func popFrom(q *[]int) (v int) {
	v = (*q)[len(*q) - 1] //取数组的最后一个元素
	*q = (*q)[:len(*q) - 1] //切片删除最后一个元素
	return
}

func (this *CQueue) DeleteHead() int {
	if len(this.q) > 0{
		//如果弹出栈有元素的话，就返回最后一个元素
		return popFrom(&this.q)
	}

	//如果弹出栈没有元素的话，就先从压入栈中的元素取出放置到弹出栈里面（这里需要注意顺序转换）
	for i := len(this.p) - 1; i >= 0; i--{
		this.q = append(this.q, this.p[i]) //压入栈的最后一个元素，到了弹出栈里面是第一个元素
	}

	//放置完毕之后，重置压入栈对象
	this.p = []int{}
	if len(this.q) == 0 {
		//如果放置完毕之后弹出栈的元素个数还是0的话，就返回-1
		return -1
	}else {
		return popFrom(&this.q)
	}
}