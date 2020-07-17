package main

import "fmt"

func main()  {
	var num = 4

	var result = Fib(num)

	fmt.Println(result)
}

func Fib(num int) int {
	if num == 1 || num == 2{
		return num
	}

	return Fib(num - 1) + Fib(num - 2)
}
