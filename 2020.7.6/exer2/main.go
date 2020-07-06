package main

import "fmt"

func main()  {
	var array = []int{1,2}

	result := uniqueOccurrences(array)

	fmt.Println(result)
}

func uniqueOccurrences(arr []int) bool {
	var tmpMap = make(map[int]int)

	for _, item := range arr{
		tmpMap[item]++
	}

	var tmpMap2 = make(map[int]int)

	for _, v := range tmpMap{
		if _, ok := tmpMap2[v]; ok{
			//如果存在
			return false
		}else {
			tmpMap2[v]++
		}
	}

	return true
}