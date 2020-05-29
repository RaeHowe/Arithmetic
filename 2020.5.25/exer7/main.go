package main

import "fmt"

func main()  {
	var str1 = "abcc"
	var str2 = "bcaa"

	result := CheckPermutation(str1, str2)
	fmt.Println(result)
}

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2){
		return false
	}

	var tmpMap = make(map[rune]int)

	for _, item := range s1{
		tmpMap[item]++
	}

	for _, item := range s2{
		tmpMap[item]--
	}

	for _, v := range tmpMap{
		if v > 0{
			return false
		}
	}

	return true
}