package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func f1(f func(x, y int) int) {
	i := f(1, 2)
	fmt.Println(i)
}

func f2(x, y int) (ret func(name string) string) {
	age := x + y
	ret = func(name string) string {
		return name
	}
	fmt.Println(age)
	fmt.Println(ret)
	return
}
func main() {

	//函数作为参数
	f1(sum)

	//函数作为返回值
	f2(2, 2)
}
