package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var str1 = "11"
	var str2 = "1"

	result := addBinary(str1, str2)

	fmt.Println(result)
}

func addBinary(a string, b string) string {
	var minIndex = minLength(a, b)

	var flagOfMin string

	if len(a) > len(b) {
		flagOfMin = "B"
	}else if len(a) < len(b){
		flagOfMin = "A"
	}



	var resultSlice []int

	//字符串补全
	if len(a) > len(b){
		tmpLength := len(a) - len(b)
		var tmpData string
		for i := 0; i < tmpLength; i++{
			tmpData += "0"
		}

		b = tmpData + b
	}else if len(a) < len(b){
		tmpLength := len(b) - len(a)
		var tmpData string
		for i := 0; i < tmpLength; i++{
			tmpData += "0"
		}

		a = tmpData + a
	}

	var flag int
	for i := 0; i < minIndex; i++{
		var tmpData int

		intA, _ := strconv.Atoi(string(a[len(a) - i - 1]))
		intB, _ := strconv.Atoi(string(b[len(b) - i - 1]))

		//判断是否需要加上进位
		if flag == 1{
			tmpData = intA + intB + 1
		}else {
			tmpData = intA + intB
		}

		var tmpRes  = tmpData % 2
		resultSlice = append(resultSlice, tmpRes)

		//重置进位标识
		if tmpData >= 2{
			flag = 1
		}else {
			flag = 0
		}
	}

	//记住要加上标识位
	if flagOfMin == "B"{
		for i := len(a) - minIndex - 1; i >= 0; i--{
			intNum, _ := strconv.Atoi(string(a[i]))
			var tmpRes int
			if flag == 1{
				tmpRes = intNum + 1
				if tmpRes >= 2{
					flag = 1
				}else {
					flag = 0
				}

				resultSlice = append(resultSlice, tmpRes % 2)
			}else {
				resultSlice = append(resultSlice, intNum)
			}
		}
	}else if flagOfMin == "A"{
		for i := len(b) - minIndex - 1; i >= 0; i--{
			intNum, _ := strconv.Atoi(string(b[i]))
			var tmpRes int
			if flag == 1{
				tmpRes = intNum + 1
				if tmpRes >= 2{
					flag = 1
				}else {
					flag = 0
				}

				resultSlice = append(resultSlice, tmpRes % 2)
			}else {
				resultSlice = append(resultSlice, intNum)
			}
		}
	}

	if flag == 1{
		resultSlice = append(resultSlice, flag)
	}

	//切片反转
	reverseArray(resultSlice)

	var result string
	for i := 0; i < len(resultSlice); i++{
		result += strconv.Itoa(resultSlice[i])
	}

	return result
}

func minLength(str1, str2 string) int {
	if len(str1) <= len(str2){
		return len(str1)
	}else {
		return len(str2)
	}
}

func reverseArray(tmpSlice []int)  {
	for i := 0; i < len(tmpSlice) / 2; i++{
		tmpSlice[i], tmpSlice[len(tmpSlice) - 1 - i] = tmpSlice[len(tmpSlice) - 1 - i], tmpSlice[i]
	}
}