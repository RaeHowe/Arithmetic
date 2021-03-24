package main

import "fmt"

func main()  {
	var array1 = []int{2, 4, 5, 8, 10}
	var array2 = []int{1, 3, 5, 7, 10}

	result := mergeArray(array1, array2)

	fmt.Println(result)
}

func mergeArray(array1, array2 []int) []int {
	var resultArray []int

	var indexA = 0
	var indexB = 0

	for indexA != len(array1) && indexB != len(array2){
		if array1[indexA] <= array2[indexB]{
			resultArray = append(resultArray, array1[indexA])
			indexA++
		}else {
			resultArray = append(resultArray, array2[indexB])
			indexB++
		}
	}

	if indexA == len(array1){
		resultArray = append(resultArray, array2[indexB:]...)
	}else if indexB == len(array2){
		resultArray = append(resultArray, array1[indexA:]...)
	}

	return resultArray
}