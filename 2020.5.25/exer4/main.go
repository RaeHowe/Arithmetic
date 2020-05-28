package main

import "fmt"

func main()  {
	var array = [][]int{{1, 2, 3}, {4, 5, 6}} //4行3列
	result := transpose(array)
	fmt.Println(result)
}

/*
	题目要求

*/

func transpose(A [][]int) [][]int {
	var c = len(A) //2
	var r = len(A[0]) //3
	var result = make([][]int, r)

	for i := 0; i < r; i++{
		for j := 0; j < c ; j++{
			result[i] = append(result[i], A[j][i])
		}
	}

	return result
}
