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
	var result1 = 1
	var result2 = 0

	var strNum = strconv.Itoa(n)

	for i := 0; i < len(strNum); i++{
		result1 = result1 * int(strNum[i] - '0')
		result2 += int(strNum[i] - '0')
	}

	return result1 - result2
}
