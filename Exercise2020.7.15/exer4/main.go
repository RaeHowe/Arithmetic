package main

import "fmt"

func main()  {
	var array = []int{-3,2,-3,4,2}

	result := minStartValue(array)

	fmt.Println(result)
}

func minStartValue(nums []int) int {
	var result = 0

	var tmpMinNum = 0

	for i := 0; i < len(nums); i++{
		tmpMinNum += nums[i]

		if tmpMinNum < result{
			result = tmpMinNum
		}
	}

	if result < 0{
		return 1 - result
	}else {
		return 1
	}
}