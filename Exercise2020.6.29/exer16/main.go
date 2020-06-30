package main

import "fmt"

func main()  {
	var array = []int{2,5,1,3,4,7}

	result := shuffle(array, 3)

	fmt.Println(result)
}

func shuffle(nums []int, n int) []int {
	var resultArray []int

	for i := 0; i < n; i++{
		resultArray = append(resultArray, nums[i])
		resultArray = append(resultArray, nums[i + n])
	}

	return resultArray
}
