package main

import "fmt"

//结构体继承

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s is move \n", a.name)
}

type dog struct {
	feet   string
	animal //只有声明成匿名结构体嵌套才可以实现继承方式
}

func main() {
	d1 := dog{
		animal: animal{name: "aa"},
		feet:   "脚",
	}
	d1.move()
	fmt.Println(d1.name)
}
