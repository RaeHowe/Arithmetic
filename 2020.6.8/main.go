package main

import "fmt"

func main()  {
	var array = []int{2, 2, 1, 1, 1, 2, 2}
	var result = majorityElement(array)
	fmt.Println(result)
}

func majorityElement(nums []int) int {
	var tmpMap = make(map[int]int)

	for i := 0; i < len(nums); i++{
		tmpMap[nums[i]]++
	}

	var maxLength = len(nums) / 2

	for k, v := range tmpMap{
		if v > maxLength{
			return k
		}
	}

	return -1
}
