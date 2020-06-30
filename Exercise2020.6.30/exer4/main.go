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
	var byteArrayNum = []byte(strNum)

	var tmpResultA = 1
	for i := 0; i < len(strNum); i++ {
		intNum, _ := strconv.Atoi(string(byteArrayNum[i]))
		tmpResultA = tmpResultA * intNum
	}

	var tmpResultB int
	for i := 0; i < len(strNum); i++{
		intNum, _ := strconv.Atoi(string(byteArrayNum[i]))
		tmpResultB += intNum
	}

	return tmpResultA - tmpResultB
}
