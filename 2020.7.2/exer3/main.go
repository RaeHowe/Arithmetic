package main

import (
	"fmt"
	"strings"
)

func main()  {
	var str1 = "waterbottle"
	var str2 = "erbottlewat"

	flag := isFlipedString(str1, str2)

	fmt.Println(flag)
}

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2){
		return false
	}

	if strings.Contains(s2 + s2, s1){
		return true
	}else {
		return false
	}
}