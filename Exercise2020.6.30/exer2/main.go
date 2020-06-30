package main

import "fmt"

func main()  {
	var array = []int{3, 10, 16, 23, 39, 40, 58, 69, 88, 90}

	result := binaryQuery(array, 90)

	fmt.Println(result)
}

func binaryQuery(tmpSlice []int, target int) int {
	var indexStart = 0
	var indexEnd = len(tmpSlice) - 1

	for indexStart <= indexEnd{
		var midIndex = (indexStart + indexEnd) / 2
		var midData = tmpSlice[midIndex]
		if midData == target{
			return midIndex
		}

		if midData > target{
			indexEnd = midIndex - 1
		}

		if midData < target{
			indexStart = midIndex + 1
		}
	}

	return -1
}
