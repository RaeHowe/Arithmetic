package main

import "fmt"

func main()  {
	var tmpArray = []int{1, 4, 3, 2}
	result := arrayPairSum(tmpArray)
	fmt.Println(result)
}

func arrayPairSum(nums []int) int {
	quickSort(nums, len(nums), 0, len(nums) - 1)

	var result = 0

	for i := 0; i < len(nums);{
		var tmpData = min(nums[i], nums[i+1])
		result += tmpData
		i += 2
	}

	return result
}

func min(a, b int) int {
	if a <= b {
		return a
	}else {
		return b
	}
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

func Partition(tmpSlice []int, length, start, end int) int{
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