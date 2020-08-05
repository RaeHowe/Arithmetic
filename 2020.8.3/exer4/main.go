package main

import "fmt"

func main()  {
	var array = []int{5,7,7,8,8,10}

	result := search(array, 8)

	fmt.Println(result)
}

func search(nums []int, target int) int {
	var start = 0
	var end = len(nums) - 1

	var resultCount = 0

	for start <= end{
		var middleIndex = (start + end) / 2

		if nums[middleIndex] == target{
			resultCount++
		}

		if nums[middleIndex] > target{
			end = middleIndex - 1
		}

		if nums[middleIndex] < target{
			start = middleIndex + 1
		}
	}

	return resultCount
}