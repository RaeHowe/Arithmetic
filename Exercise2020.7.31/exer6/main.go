package main

import "fmt"

func main()  {
	var numArray = []int{0, 5, 3, 1, 2}

	result := missingNumber(numArray)

	fmt.Println(result)
}

func missingNumber(nums []int) int {
	var length = len(nums)
	var sum = (1 + length) * length / 2

	for i := 0; i < length; i++{
		sum -= nums[i]
	}

	return sum
}