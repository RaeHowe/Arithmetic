package main

import "fmt"

func main()  {
	var num = 4
	result := fib(num)

	fmt.Println(result)
}

func fib(N int) int {
	if N == 0 || N == 1{
		return N
	}

	return fib(N - 1) + fib(N - 2)
}