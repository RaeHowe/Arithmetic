package main

import (
	"fmt"
	"strings"
)

func main()  {
	var str = "A man, a plan, a canal: Panama"

	result := isPalindrome(str)

	fmt.Println(result)
}

func isPalindrome(s string) bool {
	var newStr string

	//去除空格和符号
	for i := 0; i < len(s); i++{
		if judgementSymbol(s[i]){
			newStr += string(s[i])
		}
	}

	var indexOfStart = 0
	var indexOfEnd = len(newStr) - 1
	
	var flag = true
	for indexOfStart <= indexOfEnd{
		if strings.ToLower(string(newStr[indexOfStart])) != strings.ToLower(string(newStr[indexOfEnd])){
			flag = false
			break
		}

		indexOfStart++
		indexOfEnd--
	}

	return flag
}

func judgementSymbol(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}