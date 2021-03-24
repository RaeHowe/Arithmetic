package main

import "fmt"

func main()  {
	var array = []int{5, 1, 10, 24, 8, 7}

	bubblingSort(array)

	fmt.Println(array)
}

func bubblingSort(array []int){
	for i := 0; i < len(array); i++{
		for j := 0; j < len(array) - i - 1; j++{
			if array[j] > array[j + 1]{
				array[j], array[j + 1] = array[j + 1], array[j]
			}
		}
	}
}