package main

import "fmt"

type myInt int    //自定义类型  myint类型
type newInt = int //类型别名  类型还是int

func main() {
	var i myInt = 0
	fmt.Printf("i type is %T\n", i) //myInt
	var i2 newInt = 1
	fmt.Printf("i2 type is %T\n", i2) //int
}
