package main

import "fmt"

func main()  {
	var str = "taactcoaa"

	result := canPermutePalindrome(str)

	fmt.Println(result)
}

func canPermutePalindrome(s string) bool {
	var tmpMap = make(map[int32]int)

	for _, item := range s{
		tmpMap[item]++
	}

	var count int
	for _, v := range tmpMap{
		if v % 2 != 0{
			//不是偶数
			count++
		}
	}

	if count > 1{
		return false
	}else {
		return true
	}
}