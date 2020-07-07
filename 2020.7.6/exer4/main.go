package main

import "fmt"

func main()  {
	var array = []int{4,1,2,1,2}

	result := singleNumber(array)

	fmt.Println(result)
}

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
