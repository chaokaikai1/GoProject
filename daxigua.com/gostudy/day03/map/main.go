package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//map 是引用类型 引用类型使用前得初始化 make 来做切片和map的初始化
	var m1 map[int]string
	fmt.Println(m1 == nil)
	m1 = make(map[int]string)
	m1[0] = "aa"
	m1[1] = "bb"
	fmt.Println(m1)
	//go 语言中 ,ok ok是返回true or false前边的是返回值
	v, ok := m1[0]
	if !ok {
		println("没有该键")
	} else {
		fmt.Println(v)
	}
	// map的循环

	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//只循环k

	for k := range m1 {
		fmt.Println(k)
	}
	//只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	//删除  如果删除一个不存在的 什么都不干 no-op 不操作
	delete(m1, 0)
	fmt.Println(m1)
	//// 我们一般使用系统时间的不确定性来进行初始化 随机数初始化
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}

}
