package main

import "fmt"

func main()  {
	var array = []int{-1,3,5,1,4,2,-9}

	result := maxNonOverlapping(array, 6)

	fmt.Println(result)
}

func maxNonOverlapping(nums []int, target int) int {
	var tmpIndex = 0
	var count int

	for tmpIndex < len(nums){
		var tmpResult = target - nums[tmpIndex]
		var index = find(nums, tmpResult, tmpIndex) //如果结果在数组存在的话就删除这两个数

		if index >= 0{
			//存在
			count++
			nums = append(nums[:tmpIndex], nums[index + 1:]...)
			tmpIndex = 0
		}else {
			tmpIndex++
		}
	}

	return count
}

func find(array []int, target, tmpIndex int) int {
	for index, item := range array{
		if item == target{
			if index != tmpIndex{
				return index
			}
		}
	}

	return -1
}