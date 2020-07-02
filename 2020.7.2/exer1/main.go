package main

import (
	"fmt"
	"strings"
)

func main()  {
	var array = []string{"mass", "as", "hero", "superhero"}

	result := stringMatching(array)

	fmt.Println(result)
}

func stringMatching(words []string) []string {
	var resultMap = make(map[string]int)

	for i := 0; i < len(words); i++{
		for j := 0; j < len(words); j++{
			if i == j{
				continue
			}else {
				if strings.Contains(words[j], words[i]){
					resultMap[words[i]]++
				}
			}
		}
	}

	var resultArray []string

	for k, _ := range resultMap{
		resultArray = append(resultArray, k)
	}

	return resultArray
}
