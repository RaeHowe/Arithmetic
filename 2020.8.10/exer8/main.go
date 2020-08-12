package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	var str = "aabcccccaaa"

	result := compressString2(str)

	fmt.Println(result)
}

func compressString2(S string) string {
	if S == ""{
		return S
	}

	var sb = strings.Builder{}
	curr := S[0]
	currLen := 1

	for i := 1; i < len(S); i++{
		if S[i] == curr{
			currLen++
		}else {
			sb.WriteByte(curr)
			sb.WriteString(strconv.Itoa(currLen))
			curr = S[i]
			currLen = 1
		}
	}

	sb.WriteByte(curr)
	sb.WriteString(strconv.Itoa(currLen))

	if sb.Len() >= len(S){
		return S
	}else {
		return sb.String()
	}
}

func compressString(S string) string {
	var tmpStr string
	var tmpCount int

	var result string

	var isStart = true
	for i := 0; i < len(S); i++{
		if string(S[i]) != tmpStr{
			if isStart == false{
				result += tmpStr
				result += strconv.Itoa(tmpCount)
			}

			tmpStr = string(S[i])
			tmpCount = 1
		}else {
			tmpCount++
		}

		isStart = false
	}

	result = result + tmpStr + strconv.Itoa(tmpCount)

	if len(result) >= len(S){
		return S
	}else {
		return result
	}
}