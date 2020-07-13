package main

import "fmt"

func main()  {
	var array = []int{1,0,2,3,0,4,5,0}

	duplicateZeros(array)

	fmt.Println(array)
}

func duplicateZeros(arr []int)  {
	var length = len(arr)

	for i := 0; i < length; i++{
		if arr[i] == 0{
			//判断数组长度是否满足要求
			if i != length - 1{
				for j := length - 1; j > i; j--{
					arr[j] = arr[j - 1]
				}
			}
			//如果i是最后一个，那就不用变化

			i++
		}
	}
}
