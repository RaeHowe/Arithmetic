package main

import "fmt"

func main()  {
	var str = "leetcode"

	result := maxPower(str)

	fmt.Println(result)
}

func maxPower(s string) int {
	var tmpMaxNum = 0

	var tmpSameData string
	var count = 1
	for i := 0; i < len(s); i++{
		if string(s[i]) == tmpSameData{
			count++
		}else {
			//重置count数量
			count = 1
			tmpSameData = string(s[i])
		}

		if count > tmpMaxNum{
			tmpMaxNum = count
		}
	}

	return tmpMaxNum
}