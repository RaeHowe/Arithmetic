package main

import "fmt"

func main()  {
	var array = []int{1, 2, 2, 2, 2}
	result := minCostToMoveChips(array)

	fmt.Println(result)
}

func minCostToMoveChips(chips []int) int {
	var oddNumberCount = 0
	var evenNumberCount = 0
	for i := 0; i < len(chips); i++{
		if chips[i] % 2 == 0{
			evenNumberCount++
		}else {
			oddNumberCount++
		}
	}

	if oddNumberCount >= evenNumberCount{
		return evenNumberCount
	}else {
		return oddNumberCount
	}
}
