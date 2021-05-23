package main

import "fmt"

func variable() {
	var a int
	var b string
	fmt.Println(a, b)
}

func variable2() {
	var a int = 1
	var b string = "aa"
	fmt.Println(a, b)
	fmt.Printf("%d  %s", a, b)
}
func variable3() {
	a := 1
	b := "bb"
	fmt.Println(a, b)
}

var (
	a = 4
	b = "ckk"
	c = false
)

func variable4() {
	var (
		a = 1
		b = "aa"
		c = true
	)
	fmt.Println(a, b, c)
}
func variable5() {
	var a, b, c int = 1, 2, 3
	var d = "abc"
	fmt.Println(a, b, c)
	fmt.Println(d)
}
func main() {
	//variable4()
	//fmt.Printf("外边 a is %d b is %s  c is %v",a,b,c)
	variable5()
}
