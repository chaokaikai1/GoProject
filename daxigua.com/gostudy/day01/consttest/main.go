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
}
