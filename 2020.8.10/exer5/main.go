package main

import "fmt"

func main()  {
	var str = "PPALLL"

	result := checkRecord(str)

	fmt.Println(result)
}

func checkRecord(s string) bool {
	var countOfA int
	var countOfL int

	for _, item := range s{
		if item == 'A'{
			countOfA++
			countOfL = 0
			if countOfA > 1{
				return false
			}
		}else if item == 'L'{
			countOfL++
			if countOfL > 2{
				return false
			}
		}else if item == 'P'{
			countOfL = 0
		}
	}

	return true
}