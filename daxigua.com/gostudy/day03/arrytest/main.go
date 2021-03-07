package main

import "fmt"

func main() {
	// var a1 [3]bool
	// //初始化1
	// var a2 [2]string = [2]string{"aa", "bb"}
	// //初始化2
	// a3 := [2]int{2, 3}
	// //初始化3 根据初始值自动推算长度  用...
	// a4 := [...]int{2, 5, 6, 7, 8}
	// //初始化4 根据索引赋值
	// a5 := [4]string{0: "aa", 3: "bb"}
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)
	// fmt.Println(a4)
	// fmt.Println(a5)

	//数组的遍历
	names := [...]string{"aa", "bb", "cc"}
	cities := [...]string{"北京", "上海", "广州"}
	//
	fmt.Printf("len(names):%d ,cap(names) :%d", len(names), cap(names))
	fmt.Println("")
	//根据索引遍历
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	//for range   i,v  i是索引  也可以写哑元变量 给隐藏
	for _, v := range cities {
		fmt.Println(v)
	}

	//多维数组
	a1 := [3][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{5, 6}}
	for _, v1 := range a1 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
