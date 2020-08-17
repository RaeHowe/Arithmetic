package main

import "fmt"

func main()  {
	var str = "([)]"

	flag := isValid(str)

	fmt.Println(flag)
}

func isValid(s string) bool {
	if s == ""{
		return true
	}

	if len(s) == 1{
		return false
	}

	var stack []rune

	var tmpMap = make(map[rune]rune)

	tmpMap['('] = ')'
	tmpMap['{'] = '}'
	tmpMap['['] = ']'

	for _, item := range s{
		if _, ok := tmpMap[item]; ok{
			//是左括号
			stack = append(stack, item)
		}else {
			//是右括号
			if len(stack) == 0{
				//如果栈里面没有任何元素的话，直接返回false错误
				return false
			}
			//说明不是左括号，就跟数组的最后一个元素进行匹配，如果匹配不上，就返回false
			var flag = false
			var tmpKey rune
			for k, v := range tmpMap{
				//验证是否是标准括号
				if v == item{
					flag = true
					tmpKey = k
				}
			}

			if flag == false{
				return false
			}

			if stack[len(stack) - 1] != tmpKey{
				return false
			}else {
				stack = stack[:len(stack) - 1]
			}
		}
	}

	if len(stack) == 0{
		return true
	}else {
		return false
	}

}