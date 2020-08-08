package main

import "fmt"

func main()  {
	var str = "FlaG"

	result := detectCapitalUse(str)

	fmt.Println(result)
}

func detectCapitalUse(word string) bool {
	if 'A' <= word[0] && word[0] <= 'Z'{
		//如果首字母是大写字母，则后面字符都得是小写或者大写
		for i := 2; i < len(word); i++{
			if !((word[i] <= 'Z' && word[i - 1] <= 'Z') || (word[i] > 'Z' && word[i - 1] > 'Z')){
				return false
			}
		}
		return true
	}else if 'a' <= word[0] && word[0] <= 'z'{
		//如果首字母是小写字母，则后面字符都得是小写字母
		for i := 1; i < len(word); i++{
			if word[i] <= 'Z'{
				//有一个大写字母就报错
				return false
			}
		}

		return true
	}

	return false
}