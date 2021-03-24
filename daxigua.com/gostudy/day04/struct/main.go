package main

import "fmt"

//结构体
type person struct {
	name  string
	age   int
	hobby []string
}

func main() {
	//结构体声明
	var p person
	p.name = "aa"
	p.age = 11
	p.hobby = []string{"篮球", "足球", "乒乓球"}
	fmt.Println(p)

	//结构体初始化方式2  key value
	 
	var p2=person{
		name:"gg",
		age:22,
		hobby:[]string{"aa","bb"},
	}
	//短变量声明
	p3:=person{
		name:"aa",
		age:33,
	}
	//直接写值 按照字段顺序写
	p4:=person{
		"aa",
		22,
	}
	fmt.Println(p2)

	//匿名结构体
	var ss struct {
		name string
		age  int
	}
	ss.name = "aa"
	ss.age = 11
	fmt.Println(ss)
}
