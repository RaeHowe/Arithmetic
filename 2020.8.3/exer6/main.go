package main

import "fmt"

func main()  {
	var str1 = "abc"
	var str2 = "ahbgdc"

	flag := isSubsequence(str1, str2)

	fmt.Println(flag)
}

func isSubsequence(s string, t string) bool {
	var indexA = 0
	var indexB = 0

	var lenA = len(s)
	var lenB = len(t)

	for indexA != lenA && indexB != lenB{
		if s[indexA] == t[indexB]{
			indexA++
			indexB++
		}else {
			indexB++
		}
	}

	return indexA == lenA
}
