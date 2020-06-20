package main

import "fmt"

func main()  {
	var paths = [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}

	fmt.Println(destCity(paths))
}

func destCity(paths [][]string) string {
	var tmpMap = make(map[string]string)

	for index, _ := range paths{
		tmpMap[paths[index][0]] = paths[index][1]
	}

	var ans string
	for _, v := range tmpMap{
		if _, ok := tmpMap[v]; !ok{
			ans = v
		}
	}

	return ans
}
