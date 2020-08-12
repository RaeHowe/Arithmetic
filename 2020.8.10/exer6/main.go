package main

import (
	"fmt"
	"math"
)

func main()  {
	var str = "ab123"

	result := reformat(str)

	fmt.Println(result)
}

func reformat(s string) string {
	var numberArray []string
	var letterArray []string

	for _, item := range s{
		if 'a' <= item && item <= 'z'{
			letterArray = append(letterArray, string(item))
		}else if '0' <= item && item <= '9'{
			numberArray = append(numberArray, string(item))
		}
	}

	var tmpData = math.Abs(float64(len(letterArray) - len(numberArray)))

	if tmpData > 1{
		return ""
	}

	var indexA = 0
	var indexB = 0

	var result string
	if len(numberArray) >= len(letterArray){
		for indexA < len(letterArray) || indexB < len(numberArray){
			if indexB != len(numberArray){
				result += numberArray[indexB]
			}

			if indexA != len(letterArray){
				result += letterArray[indexA]
			}

			indexA++
			indexB++
		}
	}else {
		for indexA < len(letterArray) || indexB < len(numberArray){
			if indexA != len(letterArray){
				result += letterArray[indexA]
			}

			if indexB != len(numberArray){
				result += numberArray[indexB]
			}

			indexA++
			indexB++
		}
	}

	return result
}