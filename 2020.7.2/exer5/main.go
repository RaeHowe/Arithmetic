package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var str1 = "289"
	var str2 = "1388"

	result := sumNum(str1, str2)

	fmt.Println(result)
}

func sumNum(a, b string) string {
	result := ""
	flag := 0

	indexA, indexB := len(a) - 1, len(b) - 1

	for indexA >= 0 || indexB >= 0{
		t1, t2 := 0, 0
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

func addBinary(a string, b string) string {
	result := ""
	flag := 0       // 存储进位

	i,j := len(a)-1,len(b)-1
	for i>=0 || j>=0 {
		t1,t2 := 0,0
		if i>=0 {
			t1 = int(a[i]-'0')
		}
		if j>=0 {
			t2 = int(b[j]-'0')
		}
		sum := t1 + t2 + flag   // 计算当前位置
		switch sum {
		case 3: flag = 1; result = "1" + result
		case 2: flag = 1; result = "0" + result
		case 1: flag = 0; result = "1" + result
		case 0: flag = 0; result = "0" + result
		}
		i--
		j--
	}
	if flag == 1 {  // 最终进位
		result = "1" + result
	}
	return result
}