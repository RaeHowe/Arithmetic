package main

import "fmt"

func main()  {
	var array = []int{17,18,5,4,6,1}

	result := replaceElements(array)

	fmt.Println(result)
}

func replaceElements(arr []int) []int {
	var index = len(arr) - 1
	var resultArray = make([]int, len(arr))
	resultArray[index] = -1

	for i := index - 1; i >= 0; i--{
		resultArray[i] = maxNum(resultArray[i + 1], arr[i + 1])
	}

	return resultArray
}

func maxNum(a, b int) int {
	if a >= b {
		return a
	}else {
		return b
	}
}