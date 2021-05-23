package main

import "fmt"

//数组声明
func arry1test() {
	//只声明 默认0值
	var arry1 [3]int
	var arry2 [4]int = [4]int{1, 2, 3, 4}
	arry3 := [2]int{2, 3}
	arry4 := [...]string{"aa", "bb"} //...代表长度可变
	fmt.Println(arry1)
	fmt.Println(arry2)
	fmt.Println(arry3)
	fmt.Println(arry4)
	arry2[0] = 7
	fmt.Println(arry2)
}

//for 循环
func forArry() {
	arry := [...]string{"aa", "bb", "cc"}

	//方法1
	for i := 0; i < len(arry); i++ {
		println(arry[i])
	}
	//方法2 for range  i 代表索引  ，v代表值  也可以用_ 匿名使用
	for _, v := range arry {
		println(v)
	}
}

func main() {
	arry1test()
	//forArry()
}
