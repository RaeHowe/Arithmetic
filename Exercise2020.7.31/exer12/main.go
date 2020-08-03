package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1 = "289"
	var str2 = "11388"

	result := sumNum(str1, str2)

	fmt.Println(result)
}

func sumNum(a, b string) string {
	result := ""
	flag := 0

	indexA, indexB := len(a)-1, len(b)-1

	for indexA >= 0 || indexB >= 0 {
		t1, t2 := 0, 0
		if indexA >= 0 {
			t1 = int(a[indexA] - '0')
			indexA--
		}

		if indexB >= 0 {
			t2 = int(b[indexB] - '0')
			indexB--
		}

		sum := t1 + t2 + flag

		if sum >= 10 {
			flag = 1
		} else {
			flag = 0
		}

		result = strconv.Itoa(sum%10) + result

	}

	if flag == 1 {
		result = "1" + result
	}

	return result
}
