package main

import "fmt"

func consttest() {
	const a = 1
	const b = "abc"
	fmt.Println(a, b)
}
func consttest2() {
	const (
		a = 2
		b = "ccc"
	)
	fmt.Println(a, b)
}

//作为枚举
func consttest3() {
	const (
		java   = 1
		net    = 2
		python = 3
		golang = 4
	)
	fmt.Println(net)
}

//在常量里定义 默认从0 开始，每新增一行常量定义 加1 只要有iota出现  从iota往下的都是加1 ，iota所在的位置的值为  第一行 0  第二行1  第三行2
//如果在同一行声明多个  则还是原来的值
//_ 为匿名变量 只是站位  如果在const中声明  变量没赋值则和上边的变量同值
func consttest4() {
	const (
		java   = 1
		net    = 2
		python = iota + 1 //2+1
		golang
	)
	fmt.Println(net, python)
}
func main() {
	consttest4()
}
