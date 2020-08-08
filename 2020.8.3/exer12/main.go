package main

import "fmt"

func main()  {
	var str = "a-bC-dEf-ghIj"

	result := reverseOnlyLetters(str)

	fmt.Println(result)
}

func reverseOnlyLetters(S string) string {
	var strOfByteArr = []byte(S)

	var startIndex = 0
	var endIndex = len(strOfByteArr) - 1

	for startIndex <= endIndex{
		if !(('a' <= strOfByteArr[startIndex] && strOfByteArr[startIndex] <= 'z') || ('A' <= strOfByteArr[startIndex] && strOfByteArr[startIndex] <= 'Z')){
			startIndex++
			continue
		}

		if !(('a' <= strOfByteArr[endIndex] && strOfByteArr[endIndex] <= 'z') || ('A' <= strOfByteArr[endIndex] && strOfByteArr[endIndex] <= 'Z')){
			endIndex--
			continue
		}

		//满足要求的话,反转
		strOfByteArr[startIndex], strOfByteArr[endIndex] = strOfByteArr[endIndex], strOfByteArr[startIndex]

		startIndex++
		endIndex--
	}

	return string(strOfByteArr)
}