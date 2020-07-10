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
	strNum := strconv.Itoa(n)

	var result1 = 1 //乘积结果
	var result2 = 0 //相加结果

	for i := 0; i < len(strNum); i++{
		//获取每一位数字
		result1 = result1 * int(strNum[i] - '0')

		result2 += int(strNum[i] - '0')
	}

	return result1 - result2
}