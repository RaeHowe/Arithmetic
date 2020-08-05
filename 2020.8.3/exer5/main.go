package main

import (
	"fmt"
	"sort"
)

func main()  {
	var array1 = []int{4, 9, 5}
	var array2 = []int{9, 4, 9, 8, 4}

	result := intersect2(array1, array2)

	fmt.Println(result)
}

func intersect2(nums1 []int, nums2 []int) []int{
	var tmpMap = make(map[int]int)

	for _, item := range nums1{
		tmpMap[item]++
	}

	var resultArray []int
	for _, item := range nums2{
		if _, ok := tmpMap[item]; ok{
			//存在的话
			tmpMap[item]--
			if tmpMap[item] == 0{
				delete(tmpMap, item)
			}
			resultArray = append(resultArray, item)
		}
	}

	return resultArray
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var indexA = 0
	var indexB = 0

	var resultArray []int
	for indexA != len(nums1) && indexB != len(nums2){
		if nums1[indexA] == nums2[indexB] {
			resultArray = append(resultArray, nums1[indexA])
			indexA++
			indexB++
		}else if nums1[indexA] < nums2[indexB]{
			indexA++
		}else if nums1[indexA] > nums2[indexB]{
			indexB++
		}
	}

	return resultArray
}