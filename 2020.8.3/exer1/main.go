package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var num1 = "1189"
	var num2 = "288"

	result := addStrings(num1, num2)

	fmt.Println(result)
}

func addStrings(num1 string, num2 string) string {
	var lenA, lenB = len(num1) - 1, len(num2) - 1

	var result string
	var flag int

	for lenA >= 0 || lenB >= 0{
		var t1, t2 int

		if lenA >= 0{
			t1 = int(num1[lenA] - '0')
		}

		if lenB >= 0{
			t2 = int(num2[lenB] - '0')
		}

		var sum = t1 + t2 + flag

		if sum >= 10{
			flag = 1
		}else {
			flag = 0
		}

		result = strconv.Itoa(sum % 10) + result

		lenA--
		lenB--
	}

	if flag == 1{
		result = "1" + result
	}

	return result
}
