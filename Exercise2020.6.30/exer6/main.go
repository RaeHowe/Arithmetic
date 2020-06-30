package main

import "fmt"

func main()  {
	var num = 4

	result := fib(num)

	fmt.Println(result)
}

func fib(num int) int {
	if num == 0 || num == 1{
		return num
	}

	return fib(num - 1) + fib(num - 2)
}
