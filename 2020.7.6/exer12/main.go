package main

import (
	"fmt"
	"math"
)

func main()  {
	var array = []int{-2,1,-3,4,-1,2,1,-5,4}

	result := maxSubArray(array)

	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	var max = math.MinInt64
	var cur = 0

	for _, item := range nums{
		if cur > 0{
			cur += item
		}else {
			cur = item
		}

		if cur > max{
			max = cur
		}
	}

	return max
}