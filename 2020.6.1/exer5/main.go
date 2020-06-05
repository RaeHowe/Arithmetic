package main

import "fmt"

func main(){
	var array = []int{1, 2, 5, 9, 5, 9, 5, 5, 5}
	//var array = []int{3, 2}
	result := majorityElement(array)
	fmt.Println(result)
}

func majorityElement(nums []int) int {
	if len(nums) == 0{
		return -1
	}

	if len(nums) == 1{
		return nums[0]
	}

	if len(nums) == 2 && nums[0] != nums[1]{
		return -1
	}

	var moj = nums[0]
	var count = 1
	for i := 1; i < len(nums); i++{
		if moj == nums[i]{
			count++
		}else {
			count--
		}

		if count == 0{
			moj = nums[i]
			count++
		}
	}

	if count == 0{
		return -1
	}else {
		return moj
	}
}

func majorityElement2(nums []int) int {
	if nums == nil{
		return -1
	}

	var majorNumber = len(nums) / 2

	var tmpMap = make(map[int]int)

	for i := 0; i < len(nums); i++{
		tmpMap[nums[i]]++
	}

	for k, v := range tmpMap{
		if v > majorNumber{
			return k
		}
	}

	return -1
}