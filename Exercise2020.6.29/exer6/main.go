package main

import "fmt"

func main()  {
	var s1 = "abc"
	var s2 = "ba"

	result := CheckPermutation(s1, s2)

	fmt.Println(result)
}

func CheckPermutation(s1 string, s2 string) bool {
	var tmpMap = make(map[string]int)

	var byteArrStr = []byte(s1)
	for i := 0; i < len(byteArrStr); i++{
		tmpMap[string(byteArrStr[i])]++
	}

	var byteArrStr2 = []byte(s2)

	for i := 0; i < len(byteArrStr2); i++{
		tmpMap[string(byteArrStr2[i])]--
	}

	for _, v := range tmpMap{
		if v != 0{
			return false
		}
	}

	return true
}
