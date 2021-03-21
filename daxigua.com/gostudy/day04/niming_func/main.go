package main

import "fmt"

func main() {
	//匿名函数常声明在函数内部 没有名字的函数
	f1 := func(name string) {
		fmt.Println(name)
	}
	f1("aa")

	//匿名函数也可以直接调用执行
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)
	fmt.Println(f(5))
}

//递归
//递归一定有个退出的条件
//5 的阶乘  5*4*3*2*1
func f(n int) int {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}
