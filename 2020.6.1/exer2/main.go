package main

import "fmt"

func main(){
	var array = []int{3, 8, -10, 23, 19, -4, -14, 27}
	result := minimumAbsDifference(array)
	fmt.Println(result)
}

func minimumAbsDifference(arr []int) [][]int {
	QuickSort(arr, len(arr), 0, len(arr) - 1)

	var result = make([][]int, 0)
	var minData int //存放最小值
	var begin = true
	for i := 1; i < len(arr); i++{
		var tmpResult = arr[i] - arr[i - 1]
		if begin {
			minData = tmpResult
		}

		begin = false
		var tmpArray = make([]int, 2)
		//给数组插入数据
		tmpArray[0] = arr[i - 1]
		tmpArray[1] = arr[i]

		if tmpResult < minData{
			minData = tmpResult
			//把数组清空，并且把数组放置到二维数组里面
			result = make([][]int, 0)
			result = append(result, tmpArray)
		}else if tmpResult == minData{
			//如果和最小值相等，就直接append操作
			result = append(result, tmpArray)
		}
	}

	return result
}

func QuickSort(tmpSlice []int, length, start, end int) {
	if start == end{
		return
	}

	var index = Partition(tmpSlice, len(tmpSlice), start, end)

	if start < index{
		QuickSort(tmpSlice, len(tmpSlice), start, index - 1)
	}

	if index < end{
		QuickSort(tmpSlice, len(tmpSlice), index + 1, end)
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

	tmpSlice[small], tmpSlice[index] = tmpSlice[index], tmpSlice[small]

	return small
}
