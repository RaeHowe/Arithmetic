package main

import (
	"fmt"
	"math"
)

func main()  {
	var array = []int{-2, 1, -3,4,-1,2,1,-5,4}

	result := maxSubArray(array)

	fmt.Println(result)
}

func maxSubArray(nums []int) int{
	max := math.MinInt64

	cur := 0

	for _, num := range nums{
		if cur > 0{
			cur += num
		}else { //如果cur小于0 的话，就直接被新的num代替
			cur = num
		}

		if cur > max{
			max = cur
		}
	}

	return max
}