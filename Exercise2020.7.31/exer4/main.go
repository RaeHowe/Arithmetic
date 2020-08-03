package main

import "fmt"

func main()  {
	var nums = []int{-3, 2, -3, 4, 2}

	result := minStartValue(nums)

	fmt.Println(result)
}


func minStartValue(nums []int) int {
	var minResult = 0

	var tmp = 0
	for i := 0; i < len(nums); i++{
		tmp += nums[i]

		if tmp < minResult{
			minResult = tmp
		}
	}

	if minResult >= 1{
		return 0
	}else {
		return 1 - minResult
	}
}