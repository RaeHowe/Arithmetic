package main

import "fmt"

func main() {
	var nums = []int{0, 1, 2, 3, 4}
	var index = []int{0, 1, 2, 2, 1}
	var result = createTargetArray(nums, index)

	fmt.Println(result)
}

/*
	给你两个整数数组 nums 和 index。你需要按照以下规则创建目标数组：

目标数组 target 最初为空。
按从左到右的顺序依次读取 nums[i] 和 index[i]，在 target 数组中的下标 index[i] 处插入值 nums[i] 。
重复上一步，直到在 nums 和 index 中都没有要读取的元素。
请你返回目标数组。

题目保证数字插入位置总是存在。

 

示例 1：

输入：nums = [0,1,2,3,4], index = [0,1,2,2,1]
输出：[0,4,1,3,2]
解释：
nums       index     target
0            0        [0]
1            1        [0,1]
2            2        [0,1,2]
3            2        [0,1,3,2]
4            1        [0,4,1,3,2]
*/

func createTargetArray(nums []int, index []int) []int {
	var length = len(index)
	var target = make([]int, length)
	var countSort = 0

	for i := 0; i < length; i++{
		if index[i] == i{
			target[i] = nums[i]
		}else {
			for j := countSort; j > index[i]; j--{ // 3 3>2
				//判断移动数字的位数
				target[j] = target[j - 1]
			}

			target[index[i]] = nums[i]
		}

		countSort++
	}

	return target
}