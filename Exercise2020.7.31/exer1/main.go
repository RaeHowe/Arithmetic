package main

import "fmt"

func main()  {
	var array = []int{17, 18, 5, 4, 6, 1}

	result := replaceElements(array)

	fmt.Println(result)
}

func replaceElements(arr []int) []int {
	var resultArray = make([]int, len(arr))

	resultArray[len(arr) - 1] = -1

	for i := len(resultArray) - 2; i >= 0; i--{
		resultArray[i] = max(resultArray[i + 1], arr[i + 1])
	}

	return resultArray
}

func max(a, b int) int {
	if a >= b{
		return a
	}else {
		return b
	}
}