package main

import "fmt"

func main()  {
	var str = "loonbalxballpoon"

	result := maxNumberOfBalloons(str)

	fmt.Println(result)
}

func maxNumberOfBalloons(text string) int {
	var tmpMap = make(map[string]int)

	for _, item := range text{
		switch string(item) {
		case "b":
			tmpMap["b"]++
		case "a":
			tmpMap["a"]++
		case "l":
			tmpMap["l"]++
		case "o":
			tmpMap["o"]++
		case "n":
			tmpMap["n"]++
		}
	}

	if len(tmpMap) < 5{
		return 0
	}

	_min := 10000

	for k, v := range tmpMap{
		if k == "l" || k == "o"{
			if v / 2 < _min{
				_min = v / 2
			}
		}else {
			if v < _min{
				_min = v
			}
		}
	}

	return _min
}
