package main

import (
	"fmt"
	"strings"
)

func main()  {
	var sentence = "hellohello hellohellohello"
	var searchWord = "ell"

	result := isPrefixOfWord(sentence, searchWord)

	fmt.Println(result)
}

func isPrefixOfWord(sentence string, searchWord string) int {
	strArray := strings.Split(sentence, " ")

	var result = -1
	for i := 0; i < len(strArray); i++{
		if strings.HasPrefix(strArray[i], searchWord){
			result = i + 1
			break
		}
	}

	return result
}