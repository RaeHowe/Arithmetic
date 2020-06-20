package main

import "fmt"

func main(){
	var str = "abcdefg"

	result := reverseLeftWords(str, 2)


	fmt.Println(result)
}

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
