package main

import (
	"fmt"
	"strings"
)

func main()  {
	var str = "the sky is blue"



	result := reverseWords(str)

	fmt.Println(result)
}

func reverseWords(s string) string {
	s = strings.Trim(s, " ") //去除左右空格
	strArray := strings.Split(s, " ")

	var newArray []string

	for i := len(strArray) - 1; i >= 0; i--{
		var tmpStr = strings.TrimSpace(strArray[i])
		if len(tmpStr) > 0{
			newArray = append(newArray, strArray[i])
		}
	}

	return strings.Join(newArray, " ")
}