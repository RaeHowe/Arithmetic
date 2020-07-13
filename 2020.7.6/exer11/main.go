package main

import "fmt"

func main()  {
	var array = []int{0,1,2,2,3,0,4,2}

	result := removeElement(array, 2)

	fmt.Println(result)
}

func removeElement(nums []int, val int) int {
	var result = len(nums)
	for i := 0; i < len(nums); i++{
		if nums[i] == val{

			result--
		}
	}

	return result
}