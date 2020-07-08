package main

import "fmt"

func main()  {
	var array = []int{0,1,0,3,12}

	moveZeroes2(array)

	fmt.Println(array)
}

func moveZeroes2(nums []int)  {
	var index = 0
	for i := 0; i < len(nums); i++{
		if nums[i] != 0{
			nums[index] = nums[i]
			index++
		}
	}

	for i := index; i < len(nums); i++{
		nums[i] = 0
	}
}

func moveZeroes(nums []int)  {
	for i := 0; i < len(nums) - 1; i++{
		for j := 0; j < len(nums) - i - 1; j++{
			if nums[j] == 0{
				nums[j], nums[j + 1] = nums[j + 1], nums[j]
			}
		}
	}
}
