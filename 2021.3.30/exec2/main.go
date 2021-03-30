package main

import "fmt"

func main()  {
	var array = []int{1, 3, 3, 9, 10, 15, 9, 8, 20}

	var result = GetRepeatElement2(array)

	fmt.Println(result)
}

func GetRepeatElement(array []int) []int {
	var result []int

	var tmpHashMap = make(map[int]int)

	for i := 0; i < len(array); i++{
		if _, ok := tmpHashMap[array[i]]; ok{
			//如果取到值的话，说明重复了
			result = append(result, array[i])
		}else {
			tmpHashMap[array[i]] = 1
		}
	}

	return result
}

func GetRepeatElement2(array []int) []int {
	var result []int
	var sign [10000]bool

	for i := 0; i < len(array); i++{
		if sign[array[i]] {
			result = append(result, array[i])
		}else {
			sign[array[i]] = true
		}
	}

	return result
}