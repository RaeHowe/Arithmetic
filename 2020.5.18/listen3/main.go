package main

import "fmt"

/*
	给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。



示例 1：

输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]

示例 2：

输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]

*/

func main()  {
	var array = []int{-4, -1, 0, 3, 10}
	var result = sortedSquares(array)
	fmt.Println(result)
}

func sortedSquares(A []int) []int {
	var tmpArray = make([]int, len(A))

	for i := 0; i < len(A); i++{
		tmpArray[i] = A[i] * A[i]
	}

	quickSort(tmpArray, len(tmpArray), 0, len(tmpArray) - 1)

	return tmpArray
}

func quickSort(tmpSlice []int, length, start, end int)  {
	if start == end{
		return
	}

	var index = Partition(tmpSlice, length, start, end)

	if start < index{
		quickSort(tmpSlice, length, start, index - 1)
	}

	if index < end{
		quickSort(tmpSlice, length, index + 1, end)
	}
}

func Partition(tmpSlice []int, length, start, end int) int {
	if tmpSlice == nil || length <= 0 || start < 0 || end > length{
		return -1
	}

	var index = start
	var small = start - 1

	for ;index < end; index++{
		if tmpSlice[index] < tmpSlice[end]{
			small++
			if small != index{
				tmpSlice[small], tmpSlice[index] = tmpSlice[index], tmpSlice[small]
			}
		}
	}

	small++

	tmpSlice[small], tmpSlice[end] = tmpSlice[end], tmpSlice[small]

	return small
}