package main

import (
	"fmt"
)

func main()  {
	var array = []int{-3,2,-3,4,2}

	result := minStartValue(array)

	fmt.Println(result)
}

func minStartValue(nums []int) int {
	var tmpData = 0
	var tmpMinData = 0
	for i := 0; i < len(nums); i++{
		tmpData += nums[i]

		if tmpMinData > tmpData{
			tmpMinData = tmpData
		}
	}

	if tmpMinData >= 0{
		return 1
	}else {
		return 1 - tmpMinData
	}
}
// x + (.....) = 1
//     x = 1 - (...)
//
