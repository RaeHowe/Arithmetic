package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var num1 = "289"
	var num2 = "1388"

	result := sumNum(num1, num2)

	fmt.Println(result)
}

func sumNum(a, b string) string {
	var result = ""

	var flag = 0

	var indexA, indexB = len(a) - 1, len(b) - 1
	for indexA >= 0 || indexB >= 0{
		var t1, t2 = 0, 0
		if indexA >= 0{
			t1 = int(a[indexA] - '0')
		}

		if indexB >= 0{
			t2 = int(b[indexB] - '0')
		}

		sum := t1 + t2 + flag

		if sum >= 10 {
			flag = 1
		}else {
			flag = 0
		}

		result = strconv.Itoa(sum % 10) + result

		indexA--
		indexB--
	}

	if flag == 1{
		result = "1" + result
	}

	return result
}