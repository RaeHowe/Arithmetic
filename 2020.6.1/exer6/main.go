package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var testData = 2

	result := countLargestGroup(testData)

	fmt.Println(result)
}

func countLargestGroup(n int) int {
	if n < 10 {
		return n
	}

	var tmpMap = make(map[int]int)

	var maxValue = 0
	for i := 0; i <= n; i++{
		var tmpByteArr = []byte(strconv.Itoa(i))

		var result int
		for j := 0; j < len(tmpByteArr); j++{
			tmpInt, _ := strconv.Atoi(string(tmpByteArr[j]))
			result += tmpInt
		}

		tmpMap[result]++

		if tmpMap[result] > maxValue{
			maxValue = tmpMap[result]
		}
	}

	var result int
	for _, v := range tmpMap{
		if v == maxValue{
			result++
		}
	}

	return result
}