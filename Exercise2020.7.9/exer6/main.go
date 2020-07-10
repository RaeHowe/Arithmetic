package main

import "fmt"

func main()  {
	var array = []int{9,6,4,2,3,5,7,0,1}

	result := missingNumber(array)

	fmt.Println(result)
}

func missingNumber(nums []int) int {
	var length = len(nums)
	var sumData = (1 + length) * length / 2

	for i := 0; i < len(nums); i++{
		sumData -= nums[i]
	}

	return sumData
}
