package main

import "fmt"

func main()  {
	var num = 4

	result := fib(num)

	fmt.Println(result)
}

func fib(num int) int {
	if num == 1 || num == 2{
		return num
	}

	return fib(num - 1) + fib(num - 2)
}

func climbStairs(num int) int {
	var result = []int{0, 1, 2}
	for i := 3; i <= num; i++ {
		result = append(result, result[i-1]+result[i-2])
	}
	return result[num]
}
