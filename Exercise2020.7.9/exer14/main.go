package main

import "fmt"

func main()  {
	var array = []int{1, 8, 10, 13, 19, 22, 29}

	result := binaryQuery(array, 9)

	fmt.Println(result)


}

func binaryQuery(tmpSlice []int, target int) int {
	var start = 0
	var end = len(tmpSlice) - 1

	for start < end{
		var middleIndex = (start + end) / 2

		middleData := tmpSlice[middleIndex]
		if middleData == target{
			return middleIndex
		}

		if middleData > target{
			end = middleIndex - 1
		}

		if middleData < target{
			start = middleIndex + 1
		}
	}

	var tmpMap = make(map[int]int)

	tmpMap[1] = 1

}
