package main

import "fmt"

/*
	给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。

	换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。

	以数组形式返回答案。

	 

	示例 1：

	输入：nums = [8,1,2,2,3]
	输出：[4,0,1,1,3]
	解释：
	对于 nums[0]=8 存在四个比它小的数字：（1，2，2 和 3）。
	对于 nums[1]=1 不存在比它小的数字。
	对于 nums[2]=2 存在一个比它小的数字：（1）。
	对于 nums[3]=2 存在一个比它小的数字：（1）。
	对于 nums[4]=3 存在三个比它小的数字：（1，2 和 2）。
*/

func main()  {
	var array = []int{8, 1, 2, 2, 3}
	result := smallerNumbersThanCurrent(array)
	fmt.Println(result)
}

func smallerNumbersThanCurrent(nums []int) []int {
	//维持一个桶，使用这个桶来维持nums数组中每个元素的索引位置
	var bucket [101]int
	var length = len(nums)

	//让nums里面的元素，对应到桶中的位置元素+1
	for i := 0; i < length; i++{
		bucket[nums[i]]++ //使用++而不是 =1 的原因是在nums数组中，可能存在重复值的情况
	}

	for i := 1; i < len(bucket); i++{
		bucket[i] += bucket[i - 1]
	}

	var result = make([]int, length)
	for i := 0; i < len(result); i++{
		result[i] = bucket[nums[i] - 1]
	}

	return result
}