package main

import "fmt"

func main()  {
	var array = []int{2, 3, 1, 0, 2, 5, 3}
	result := findRepeatNumber(array)
	fmt.Println(result)
}

func findRepeatNumber(nums []int) int {
	var tmpMap = make(map[int]int)

	for i := 0; i < len(nums); i++{
		if tmpMap[nums[i]] == 1{
			return nums[i]
		}else {
			tmpMap[nums[i]] = 1
		}
	}

	return -1
}