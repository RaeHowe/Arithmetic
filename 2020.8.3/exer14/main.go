package main

import "fmt"

func main()  {
	var str = "abBAcC"

	result := makeGood(str)

	fmt.Println(result)
}

func makeGood(s string) string {
	var startIndex int
	var endIndex = startIndex + 1

	for endIndex < len(s){
		if s[startIndex] == s[endIndex] - 32 || s[startIndex] == s[endIndex] + 32{
			//说明相邻的是大小写字母，需要从字符串中删除
			s = s[:startIndex] + s[endIndex + 1:]
			startIndex = 0
			endIndex = startIndex + 1

			continue
		}

		startIndex++
		endIndex++
	}

	return s
}