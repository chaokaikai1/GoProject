package main

import "fmt"

func main() {
	//指针  & 取地址  * 取指针保存的原始的变量值
	var a int = 3
	b := &a
	c := *b
	fmt.Println(b)
	fmt.Println(c)

	//空指针不能赋值
	// var a1 *int
	// *a1 = 100
	// fmt.Printf("a1 %T ,a1=%v", a1, *a1)

	//new 函数 只传递一个参数 返回该类型的指针
	//new 能用用于基础数据的分配内存
	var a1 = new(int)
	*a1 = 100
	fmt.Printf("a1 type is  %T", a1) //*int
	fmt.Println()
	fmt.Printf("&a1=%d  *a1=%d", &a1, *a1) //&a1=824633745456  *a1=100

	//make 只能用于  切片slice  map chan   这些的分配内存 （new 基础类型的分配内存）
	//返回的是类型本身 而不是类型的指针（区别于new,new 返回是该类型的指针）

}
