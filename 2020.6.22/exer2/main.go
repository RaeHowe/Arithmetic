package main

import (
	"fmt"
	"strings"
)

func main(){
	var str = "Let's take LeetCode contest"

	result := reverseWords(str)
	fmt.Println(result)
}

func reverseWords(s string) string {
	var tmpArray = strings.Split(s, " ")

	for i := 0; i < len(tmpArray); i++{
		var tmpItemArray = []byte(tmpArray[i])

		for j := 0; j < len(tmpItemArray) / 2; j++{
			tmpItemArray[j], tmpItemArray[len(tmpItemArray) - j - 1] = tmpItemArray[len(tmpItemArray) - j - 1], tmpItemArray[j]
		}

		tmpArray[i] = string(tmpItemArray)
	}

	var result string
	for i := 0; i < len(tmpArray); i++{
		result += tmpArray[i]

		if i != len(tmpArray) - 1{
			result += " "
		}
	}

	return result
}
