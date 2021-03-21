package main

import "fmt"

//结构体
type person struct {
	name  string
	age   int
	hobby []string
}

func main() {
	var p person
	p.name = "aa"
	p.age = 11
	p.hobby = []string{"篮球", "足球", "乒乓球"}
	fmt.Println(p)
	//匿名结构体
	var ss struct {
		name string
		age  int
	}
	ss.name = "aa"
	ss.age = 11
	fmt.Println(ss)
}
