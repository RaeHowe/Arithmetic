package main

import "fmt"

func main()  {
	var nums1 = []int{4,9,5}
	var nums2 = []int{9,4,9,8,4}

	result := intersection(nums1, nums2)

	fmt.Println(result)
}

func intersection(nums1 []int, nums2 []int) []int {
	var tmpMap1 = make(map[int]int)

	for _, item := range nums1{
		if _, ok := tmpMap1[item]; !ok{
			tmpMap1[item]++
		}
	}

	var tmpMap2 = make(map[int]int)

	for _, item := range nums2{
		if _, ok := tmpMap2[item]; !ok{
			tmpMap2[item]++
		}
	}

	var result []int
	for k, _ := range tmpMap2{
		if _, ok := tmpMap1[k]; ok{
			result = append(result, k)
		}
	}

	return result
}
