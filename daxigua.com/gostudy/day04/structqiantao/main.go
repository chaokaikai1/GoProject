package main

import "fmt"

type person struct {
	name string
	age  int
	addr address
}
type student struct {
	name    string
	sno     string
	address //直接声明类型  匿名字段 也是匿名嵌套结构体
}
type address struct {
	province string
	city     string
}

//结构体嵌套
func main() {
	p1 := person{
		name: "aa",
		age:  11,
		addr: address{
			province: "河南",
			city:     "焦作",
		},
	}

	s1 := student{
		name: "bb",
		address: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.addr.province)
	fmt.Println(s1.city) //匿名嵌套结构体可以直接访问嵌套的结构体字段
}
