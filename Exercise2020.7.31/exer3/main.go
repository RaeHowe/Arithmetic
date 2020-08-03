package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var nums = []int{12, 345, 2, 6, 7896}

	result := findNumbers(nums)

	fmt.Println(result)
}


func findNumbers(nums []int) int {
	var result = 0

	for i := 0; i < len(nums); i++{
		var strNum = strconv.Itoa(nums[i])
		if len(strNum) % 2 == 0{
			result++
		}
	}

	return result
}