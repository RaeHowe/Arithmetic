package main

import "fmt"

func main()  {
	var str = "abb"

	result := removePalindromeSub(str)

	fmt.Println(result)
}

func removePalindromeSub(s string) int {
	if s == ""{
		return 0
	}

	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		if s[i] != s[j]{
			return 2
		}
	}

	return 1
}