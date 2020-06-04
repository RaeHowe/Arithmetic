package main

import "fmt"

func main(){
	var array = []int{0, 2, 3, 4, 5}
	result := findMagicIndex(array)
	fmt.Println(result)
}

func findMagicIndex(nums []int) int {
	for index, value := range nums{
		if index == value{
			return value
		}
	}

	return -1
}