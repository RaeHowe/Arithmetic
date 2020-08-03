package main

import "fmt"

func main()  {
	var num = 5

	result := fib(num)

	fmt.Println(result)
}

func fib(num int) int {
	var numArray = []int{0, 1, 2}

	for i := 3; i <= num; i++{
		numArray = append(numArray, numArray[i - 1] + numArray[i - 2])
	}

	return numArray[num]
}