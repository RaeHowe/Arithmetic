package main

import "fmt"

func main()  {
	var s1 = "abc"
	var s2 = "bad"

	result := CheckPermutation(s1, s2)

	fmt.Println(result)
}

func CheckPermutation(s1 string, s2 string) bool {
	var tmpMap = make(map[string]int)

	for i := 0; i < len(s1); i++{
		tmpMap[string(s1[i])]++
	}

	for i := 0; i < len(s2); i++{
		if _, ok := tmpMap[string(s2[i])]; ok{
			tmpMap[string(s2[i])]--
		}
	}

	for _, v := range tmpMap{
		if v != 0{
			return false
		}
	}

	return true
}
