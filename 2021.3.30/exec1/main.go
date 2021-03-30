package main

import "fmt"

func main()  {
	var array = []int{10, 2, 28, 9, 14, 1, 7}

	quickSort(array, len(array), 0, len(array) - 1)

	fmt.Println(array)
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
	var small = index - 1

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