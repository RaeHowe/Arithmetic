package main

import "fmt"

func main()  {
	result := canConstruct("aa", "ab")

	fmt.Println(result)
}

func canConstruct(ransomNote string, magazine string) bool {
	var tmpMap = make(map[int32]int)

	for _, item := range magazine{
		tmpMap[item]++
	}

	for _, item := range ransomNote{
		if _, ok := tmpMap[item]; ok{
			tmpMap[item]--
			if tmpMap[item] == 0{
				delete(tmpMap, item)
			}
		}else {
			return false
		}
	}

	return true
}