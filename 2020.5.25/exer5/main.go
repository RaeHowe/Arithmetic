package main

import "fmt"

func main()  {
	var array = []int{4, 2, 6, 7, 1, 3}
	result := sortArrayByParityII(array)
	fmt.Println(result)
}

func sortArrayByParityII(A []int) []int {
	var resultArray []int
	var oddNumberArray []int
	var sacnNumberArray []int

	for i := 0; i < len(A); i++{
		if A[i] % 2 == 0{
			//偶数放入到偶数数组里面
			oddNumberArray = append(oddNumberArray, A[i])
		}else {
			sacnNumberArray = append(sacnNumberArray, A[i])
		}
	}

	var indexOfOdd = 0
	var indexOfScan = 0
	var resultArrayIndex = 0

	for resultArrayIndex != len(A){
		if resultArrayIndex % 2 == 0{
			//放入偶数
			resultArray = append(resultArray, oddNumberArray[indexOfOdd])
			indexOfOdd++
			resultArrayIndex++
		}else {
			//放入奇数
			resultArray = append(resultArray, sacnNumberArray[indexOfScan])
			indexOfScan++
			resultArrayIndex++
		}
	}

	return resultArray
}
