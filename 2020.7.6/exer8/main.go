package main

import "fmt"

func main()  {
	var array = []int{1,2,2,6,6,6,6,7,10}

	result := findSpecialInteger(array)

	fmt.Println(result)
}

func findSpecialInteger(arr []int) int {
	var length = len(arr)

	var tmpMap = make(map[int]int)

	for _, item := range arr{
		tmpMap[item]++
	}

	for k, v := range tmpMap{
		if v > length / 4{
			return k
		}
	}

	return -1
}