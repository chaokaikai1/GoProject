package main

import "fmt"

//普通循环
func for1() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//死循环
func for2() {
	for {
		fmt.Println("abc")
	}
}

//数组的for  range到数组在做
func main() {
	for1()
}
