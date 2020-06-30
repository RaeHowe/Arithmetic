package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var array = []int{12, 345, 2, 6, 7896}

	result := findNumbers(array)

	fmt.Println(result)
}

func findNumbers(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++{
		var byteArrNum = []byte(strconv.Itoa(nums[i]))
		var tmpResult int
		for j := 0; j < len(byteArrNum); j++{
			tmpResult++
		}

		if tmpResult % 2 == 0{
			result++
		}
	}

	return result
}
