package main

import "fmt"

func main() {
	const name = "name"
	const addr, age, isman = "aaa", 1, true
	fmt.Println(name)
	fmt.Printf("addr is %s ,age is %d ,isman is %t", addr, age, isman)
	fmt.Println()
	//常量用作枚举
	const (
		Man   = 0
		Women = 1
	)
	fmt.Println(Man)

	//iota  特殊常量 可以被编辑器修改的常量
	//在常量里定义 默认从0 开始，每新增一行常量定义 加1
	//如果在同一行声明多个  则还是原来的值
	//_ 为匿名变量 只是站位  如果在const中声明  变量没赋值则和上边的变量同值
	const (
		a1 = iota
		a2 = iota
		_
		z
		a3 = iota
	)

	const (
		b1, b2 = iota + 1, iota + 2
		b3, b4 = iota + 2, iota + 3
	)
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(z)
}
