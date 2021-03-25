package main

import "fmt"

func main()  {
	var array = []int{1, 4, 10, 13, 18, 20, 29, 30, 39, 51, 60}

	var index = binarySearch(array, 55)

	fmt.Println(index)
}

func binarySearch(array []int, target int) int {
	var middleIndex = len(array) / 2

	var startIndex = 0
	var endIndex = len(array) - 1

	for startIndex <= endIndex{
		if array[middleIndex] == target{
			return middleIndex
		}

		if array[middleIndex] > target{
			endIndex = middleIndex - 1
		}

		if array[middleIndex] < target{
			startIndex = middleIndex + 1
		}

		middleIndex = (startIndex + endIndex) / 2
	}

	return -1
}
