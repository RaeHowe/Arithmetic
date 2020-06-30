package main

import "fmt"

func main()  {
	var array = []int{9,6,4,2,3,5,7,0,1}

	result := missingNumber(array)

	fmt.Println(result)
}

func missingNumber(nums []int) int {
	var length = len(nums)

	var sum =  (1 + len(nums)) * 9 / 2

	for i := 0; i < length; i++{
		sum = sum - nums[i]
	}

	return sum
}