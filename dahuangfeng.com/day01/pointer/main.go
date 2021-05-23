package main

import "fmt"

func val(a int) {
	a++
}

func p_val(a *int) {
	*a++
}

func main() {
	//指针 就是指向内存地址
	//值传递 copy一份值 不会改变原始内存上的值
	var a int = 1
	val(a)
	fmt.Printf("a val is %d\n", a)

	//引用传递 传递地址 修改内存地址上的值
	p_val(&a)
	fmt.Printf("a p_val is %d\n", a)

}
