package main

import "fmt"

/*
	实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：

输入: s = "leetcode"
输出: false

示例 2：

输入: s = "abc"
输出: true
*/

func main()  {
	var flag = isUnique("helo")
	fmt.Println(flag)
}

func isUnique(astr string) bool {
	var tmpMap = make(map[byte]int)

	var tmpByteArr = []byte(astr)
	for _, item := range tmpByteArr{
		if tmpMap[item] == 1{
			return false
		}else {
			tmpMap[item] = 1
		}
	}

	return true
}