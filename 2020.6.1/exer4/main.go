package main

import "fmt"

func main()  {
	var array = []int{9,6,4,2,3,5,7,0,1}
	res := missingNumber(array)
	fmt.Println(res)
}

/*
数组nums包含从0到n的所有整数，但其中缺了一个。请编写代码找出那个缺失的整数。你有办法在O(n)时间内完成吗？

示例 1：
输入：[3,0,1]
输出：2

示例 2：
输入：[9,6,4,2,3,5,7,0,1]
输出：8
*/

func missingNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++{
		res = res + i + 1 -nums[i] //
	}

	return res
}