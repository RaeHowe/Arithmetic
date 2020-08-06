package main

import (
	"fmt"
	"strings"
)

func main()  {
	var array = []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}

	result := numUniqueEmails(array)

	fmt.Println(result)
}

func numUniqueEmails(emails []string) int {
	res := make(map[string]bool)
	for _,value := range emails{
		s:= simplify(value)
		res[s] = true
	}
	return len(res)
}

func simplify(src string)string{
	srcaite := strings.Split(src,"@")
	srcaite[0] = strings.Replace(srcaite[0],".","",-1)
	srcs := strings.Split(srcaite[0],"+")
	return srcs[0] +"@"+ srcaite[1]
}
