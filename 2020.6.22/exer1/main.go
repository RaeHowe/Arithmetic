package main

import "fmt"

func main()  {
	var str = "RLRRLLRLRL"

	result := balancedStringSplit(str)

	fmt.Println(result)
}

func balancedStringSplit(s string) int {
	l, r := 0, 0
	var result = 0

	for i := 0; i < len(s); i++{
		if s[i] == 'L'{
			l++
		}

		if s[i] == 'R'{
			r++
		}

		if l == r{
			l = 0
			r = 0
			result++
		}
	}

	return result
}
