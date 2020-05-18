package main

import "fmt"

/*
给你一个整数 n，请你返回 任意 一个由 n 个 各不相同 的整数组成的数组，并且这 n 个数相加和为 0 。



示例 1：

输入：n = 5
输出：[-7,-1,1,3,4]
解释：这些数组也是正确的 [-5,-1,1,2,3]，[-3,-1,2,-2,4]。
*/

func main()  {
	result := sumZero(5)
	fmt.Println(result)
}

func sumZero(n int) []int {
	var target []int

	for i := 1; i <= n / 2; i++{
		target = append(target, i)
		target = append(target, -i)
	}

	if n % 2 != 0 {
		target = append(target, 0)
	}

	return target
}
