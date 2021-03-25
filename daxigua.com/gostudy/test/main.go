package main

import "fmt"

type dog struct {
	name string
	age  int
}

func (d dog) eat() {
	fmt.Printf("dog name is  %s  ,say wangwang", d.name)
}
func main() {
	d1 := dog{
		name: "aa",
	}
	d1.eat()
}
