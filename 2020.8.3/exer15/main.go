package main

import (
	"fmt"
	"strconv"
)

func main()  {
	result := findKthBit(1, 1)

	fmt.Println(int(result) - '0')
}

func findKthBit(n int, k int) byte {
	if n == 1{
		return '0'
	}

	var tmpStrResult = strconv.Itoa(0)

	for n > 1{
		tmpStrResult = tmpStrResult + "1" + reverse(revert(tmpStrResult))

		n--
	}

	fmt.Println(tmpStrResult)

	return tmpStrResult[k - 1]
}

func revert(str string) string {
	var tmpByteArr = []byte(str)
	for index, item := range tmpByteArr{
		if item == '0'{
			tmpByteArr[index] = '1'
		}else if item == '1'{
			tmpByteArr[index] = '0'
		}
	}

	return string(tmpByteArr)
}

func reverse(str string) string {
	var tmpByteArr = []byte(str)

	for i := 0; i < len(tmpByteArr) / 2; i++{
		tmpByteArr[i], tmpByteArr[len(tmpByteArr) - 1 - i] = tmpByteArr[len(tmpByteArr) - 1 - i], tmpByteArr[i]
	}

	return string(tmpByteArr)
}