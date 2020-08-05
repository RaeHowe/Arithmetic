package main

import "fmt"

func main()  {
	var array = []int{5,7,7,8,8,8,8,9,10}

	result := search(array, 8)

	fmt.Println(result)
}

func search(nums []int, target int) int {
	var start = 0
	var end = len(nums) - 1

	//搜索右边界
	for start <= end{
		var middle = (start + end) / 2

		if nums[middle] <= target{
			start = middle + 1
		}else {
			end = middle - 1
		}
	}

	var right = start
	//如果数组中没有target，提前返回

	if end >= 0 && nums[end] != target{
		return 0
	}

	//搜索左边界
	start = 0
	end = len(nums) - 1

	for start <= end{
		var middle = (start + end) / 2

		if nums[middle] < target{
			start = middle + 1
		}else {
			end = middle - 1
		}
	}

	var left = end

	return right - left - 1
}