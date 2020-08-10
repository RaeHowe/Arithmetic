package main

import "fmt"

func main()  {
	var str = "00110011"

	result := countBinarySubstrings(str)

	fmt.Println(result)
}

func countBinarySubstrings(s string) int {
	var countArray []int
	var ptr = 0
	var length = len(s)

	for ptr < length{
		var tmpStartNum = s[ptr]
		var count int

		for ptr < length && s[ptr] == tmpStartNum{
			count++
			ptr++
		}

		countArray = append(countArray, count)
	}

	var ans = 0

	for i := 1; i < len(countArray); i++{
		ans += min(countArray[i], countArray[i - 1])
	}

	return ans
}

func min(a, b int) int {
	if a < b{
		return a
	}

	return b
}
