package main

import (
	"fmt"
)

func main() {
	//数组必须指定长度和类型
	//长度也是数组声明的一部分[2]string [3]string  是两个类型
	//不确定数组长度 用...
	//nameArry:=[]string{"aa","bb"}这种声明和[...]string{"aa","bb"} 类似 但是他是切片
	nameArry := [2]string{"aa", "bb"}
	var ageArry [2]int = [2]int{1, 2}
	var balance = [...]float32{1.0, 2.0}
	balance2 := [2]string{"ccc", "dddd"}
	//根据索引赋值  0  4 是索引  未赋值的默认值是0
	arry2 := [5]int{0: 1, 4: 2}
	fmt.Println(arry2)
	fmt.Println(ageArry[0])
	fmt.Println(nameArry[0])
	fmt.Println(balance[0])
	fmt.Println(balance2[0])

	//数组的遍历
	//根据索引遍历
	for i := 0; i < len(nameArry); i++ {
		fmt.Println(nameArry[i])
	}
	//for range
	for i, v := range balance2 {
		fmt.Println(i, v)
	}

}
