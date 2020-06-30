package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var num = 234

	result := subtractProductAndSum(num)

	fmt.Println(result)
}

func subtractProductAndSum(n int) int {
	var strNum = strconv.Itoa(n)
	var arrayNum = []byte(strNum)

	var result1 = 1
	var result2 = 0

	for i := 0; i < len(arrayNum); i++{
		tmpInt, _ := strconv.Atoi(string(arrayNum[i]))
		result1 = result1 * tmpInt
	}

	for i := 0; i < len(arrayNum); i++{
		tmpInt, _ := strconv.Atoi(string(arrayNum[i]))
		result2 += tmpInt
	}

	return result1 - result2
}
