package main

import (
	"fmt"
	"strings"
)

func main()  {
	var word = []string{"cat", "bt", "hat", "tree"}
	var chars = "atach"

	result := countCharacters(word, chars)
	fmt.Println(result)
}

func countCharacters(words []string, chars string) int {
	var result = 0

	loop: for _, item := range words{
		for _, char := range item{
			if strings.Count(item, string(char)) > strings.Count(chars, string(char)){
				continue loop
			}
		}

		result += len(item)
	}

	return result
}