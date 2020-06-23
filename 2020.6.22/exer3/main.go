package main

import (
	"fmt"
	"math"
)

func main()  {
	var strA = "aba"
	var strB = "cdc"

	result := findLUSlength(strA, strB)

	fmt.Println(result)
}

func findLUSlength(a string, b string) int {
	if a == b{
		return -1
	}

	return int(math.Max(float64(len(a)), float64(len(b))))
}
