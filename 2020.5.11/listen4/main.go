package main

import (
	"fmt"
)

func main() {
	var array = []int{17,18,5,4,6,1}
	result := replaceElements(array)
	fmt.Println(result)
}

func replaceElements(arr []int) []int {
	var length = len(arr)
	var result = make([]int, length)

	result[length - 1] = -1
	for i := length - 2; i >= 0; i--{
		result[i] = MaxNum(result[i + 1], arr[i + 1])
	}

	return result
}

func MaxNum(a, b int) int {
	if a >= b{
		return a
	}else {
		return b
	}
}