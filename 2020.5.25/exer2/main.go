package main

import "fmt"

func main(){
	var array = []int{4,3,2,2,4,1,3,4,3}
	result := findLucky(array)
	fmt.Println(result)
}

func findLucky(arr []int) int{
	var result = -1
	var tmpMap = make(map[int]int)

	var tmpSlice []int

	for i := 0; i < len(arr); i++{
		if tmpMap[arr[i]] == 0{
			//map中没有对应的元素
			tmpMap[arr[i]] = 1
			tmpSlice = append(tmpSlice, arr[i])
		}else {
			tmpMap[arr[i]] += 1
		}
	}

	//QuickSort(tmpSlice, len(tmpSlice), 0, len(tmpSlice) - 1)
	BubblingSort(tmpSlice)

	for i := len(tmpSlice) - 1; i >= 0; i--{
		if tmpMap[tmpSlice[i]] == tmpSlice[i]{
			result = tmpSlice[i]
			break
		}
	}

	return result
}

func BubblingSort(tmpSlice []int)  {
	for i := 0; i < len(tmpSlice) - 1; i++{
		for j := 0; j < len(tmpSlice) - i - 1; j++{
			if tmpSlice[j] > tmpSlice[j + 1]{
				tmpSlice[j], tmpSlice[j + 1] = tmpSlice[j + 1], tmpSlice[j]
			}
		}
	}
}

func QuickSort(tmpSlice []int, length, start, end int)  {
	if start == end{
		return
	}

	var index = Partition(tmpSlice, length, start, end)

	if start < index{
		QuickSort(tmpSlice, length, start, index - 1)
	}

	if index < end{
		QuickSort(tmpSlice, length, index + 1, end)
	}
}

func Partition(tmpSlice []int, length, start, end int) int {
	if tmpSlice == nil || length <= 0 || start < 0 || end > length{
		return -1
	}

	var index = start
	var small = start - 1

	for ; index < end; index++{
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