package main

import "fmt"

func main(){
	var array = []int{1,2,3,4}

	result := runningSum(array)

	fmt.Println(result)
}

func runningSum(nums []int) []int {
	var result []int

	var tmpSum = 0
	for _, item := range nums{
		tmpSum += item
		result = append(result, tmpSum)
	}

	return result
}