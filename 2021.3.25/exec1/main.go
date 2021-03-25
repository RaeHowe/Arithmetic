package main

import "fmt"

func main()  {
	var array = []int{6, 1, 10, 9, 14, 2, 5}

	quickSort(array, len(array), 0, len(array) - 1)

	fmt.Println(array)
}

func quickSort(array []int, length, start, end int)  {
	if start == end{
		return
	}

	var index = Partition(array, length, start, end)

	if start < index{
		quickSort(array, length, start, index - 1)
	}

	if index < end{
		quickSort(array, length, index + 1, end)
	}
}

func Partition(array []int, length, start, end int) int {
	if array == nil || length == 0 || start < 0 || end > length{
		return -1
	}

	var index = start
	var small = index - 1

	for ;index < end; index++{
		if array[index] < array[end]{
			small++
			if small != index{
				array[small], array[index] = array[index], array[small]
			}
		}
	}

	small++

	array[small], array[end] = array[end], array[small]

	return small
}
