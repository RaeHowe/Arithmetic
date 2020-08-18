package main

import (
	"fmt"
	"strings"
)

func main()  {
	var array = []string{"bella","label","roller"}

	result := commonChars(array)

	fmt.Println(result)
}

func commonChars(A []string) []string {
	var mid = make(map[string]int)

	for _, item := range A[0]{
		mid[string(item)]++
	}

	var res []string
	for j := 1 ; j < len(A) ; j++{
		for k, v := range mid{
			if strings.Count(A[j], k) <= v {
				mid[k] = strings.Count(A[j],k)
			}
		}
	}
	for k, v := range mid{
		for i := 1 ; i <= v ; i++{
			res = append(res, k)
		}
	}

	return  res
}