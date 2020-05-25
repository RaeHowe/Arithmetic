package main

import "fmt"

func main()  {
	var array = []int{3, 1, 2, 4}
	result := sortArrayByParity(array)
	fmt.Println(result)
}

func sortArrayByParity(A []int) []int {
	var a = 0

	for ; a < len(A); a++{
		var b = len(A) - 1
		if a == b{
			break
		}
		for A[a] % 2 != 0{
			//奇数，和最后一个元素进行交换
			A[a], A[b] = A[b], A[a]
			if a < b{
				b--
			}else {
				break
			}
		}
	}

	return A
}
