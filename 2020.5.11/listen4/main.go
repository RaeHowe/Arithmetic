package main

import (
	"fmt"
)

func main() {
	var array = []int{17,18,5,4,6,1}
	result := replaceElements(array)
	fmt.Println(result)
}

/*
	题目要求

给你一个数组 arr ，请你将每个元素用它右边最大的元素替换，如果是最后一个元素，用 -1 替换。

完成所有替换操作后，请你返回这个数组。



示例：

输入：arr = [17,18,5,4,6,1]
输出：[18,6,6,6,1,-1]
*/

func replaceElements(arr []int) []int {
	var length = len(arr)
	var result = make([]int, length)

	result[length - 1] = -1
	for i := length - 2; i >= 0; i--{
		result[i] = MaxNum(result[i + 1], arr[i + 1])
	}

	return result
}

func MaxNum(a, b int) int {
	if a >= b{
		return a
	}else {
		return b
	}
}