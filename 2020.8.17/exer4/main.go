package main

import "fmt"

func main()  {
	var array = []int{2,1,2,5,3,2}

	result := repeatedNTimes(array)

	fmt.Println(result)
}

func repeatedNTimes(A []int) int {
	var n = len(A) / 2

	var tmpMap = make(map[int]int)

	for _, item := range A{
		tmpMap[item]++
	}

	for k, v := range tmpMap{
		if v == n{
			return k
		}
	}

	return -1
}