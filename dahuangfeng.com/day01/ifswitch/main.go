package main

import "fmt"

//传递a参数  返回string
func iftest(a int) string {
	if a < 10 {
		return "<"
	} else if a == 10 {
		return "="
	} else {
		return ">"
	}
}

//switch test
//swith 没有  括号  也没有break  语句
func switchtest(a, b int, op string) int {
	result := 0
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	}
	return result
}
func main() {
	//str:=iftest(10)
	//fmt.Println(str)
	result := switchtest(2, 1, "+")
	fmt.Println(result)
}
