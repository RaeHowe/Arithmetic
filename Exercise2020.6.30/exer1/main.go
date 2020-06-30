package main

import "fmt"

func main()  {
	var arr = []int{17, 18, 5, 4, 6, 1}
	result := replaceElements(arr)

	fmt.Println(result)
}

func replaceElements(arr []int) []int {
	var length = len(arr)
	var resultArr = make([]int, length)
	resultArr[length - 1] = -1

	for i := length - 2; i >= 0; i--{
		resultArr[i] = maxNum(resultArr[i + 1], arr[i + 1])
	}

	return resultArr
}

func maxNum(a, b int) int {
	if a >= b{
		return a
	}else {
		return b
	}
}
