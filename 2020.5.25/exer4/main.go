package main

import "fmt"

func main()  {
	var array = [][]int{{1, 2, 3}, {4, 5, 6}} //4行3列
	result := transpose(array)
	fmt.Println(result)
}

/*
	题目要求
	给定一个矩阵 A， 返回 A 的转置矩阵。

矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。



示例 1：

输入：[[1,2,3],[4,5,6],[7,8,9]]
输出：[[1,4,7],[2,5,8],[3,6,9]]

示例 2：

输入：[[1,2,3],[4,5,6]]
输出：[[1,4],[2,5],[3,6]]
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
