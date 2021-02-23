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

	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
