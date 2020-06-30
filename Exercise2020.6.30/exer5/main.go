package main

import "fmt"

func main()  {
	var array = []int{-3,2,-3,4,2}

	result := minStartValue(array)

	fmt.Println(result)

}

func minStartValue(nums []int) int {
	var tmpMinData = 0

	var tmpSum = 0
	for i := 0; i < len(nums); i++{
		tmpSum += nums[i]

		if tmpSum < tmpMinData{
			tmpMinData = tmpSum
		}
	}

	if tmpMinData < 0{
		return 1 - tmpMinData
	}else {
		return 1
	}
}
