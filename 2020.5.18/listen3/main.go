package main

import "fmt"

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