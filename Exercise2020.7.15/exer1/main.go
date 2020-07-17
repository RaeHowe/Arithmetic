package main

import "fmt"

func main()  {
	var array = []int{17,18,5,4,6,1}

	result := replaceElements(array)

	fmt.Println(result)
}

func replaceElements(arr []int) []int {
	var length = len(arr)
	var resultArray = make([]int, length)

	resultArray[length - 1] = -1

	for i := length - 1; i > 0; i--{
		resultArray[i - 1] = maxNum(arr[i], resultArray[i])
	}

	return resultArray
}

func maxNum(a, b int) int {
	if a >= b{
		return a
	}else {
		return b
	}
}