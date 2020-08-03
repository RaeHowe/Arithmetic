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

	var resultOfReduce = 0
	var resultOfProduct = 1

	for i := 0; i < len(strNum); i++{
		resultOfReduce += int(strNum[i] - '0')
		resultOfProduct = resultOfProduct * int(strNum[i] - '0')
	}

	return resultOfProduct - resultOfReduce
}