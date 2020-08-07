package main

import (
	"fmt"
	"strings"
)

func main()  {
	//var str = "Mr John Smith  "
	var str = "                     "
	var length = 5

	result := replaceSpaces(str, length)

	fmt.Println(result)
}


func replaceSpaces(S string, length int) string {
	return strings.ReplaceAll(S[:length]," ", "%20")
}