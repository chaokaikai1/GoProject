package main

import (
	"fmt"
	"time"
)

//goroutine 并发 开启线程  在调用方法前添加个go

func hello(i int) {
	fmt.Printf("hello %v \n", i)
}

func main() {
	go hello(777)
	for i := 1; i < 100; i++ {
		//go hello(i) //开启一个goroutine 去执行
		//用匿名函数写
		go func(i int) {
			fmt.Println(i)
		}(i) //匿名函数 写成闭包的形式 要不i会输出重复的值
	}
	fmt.Println("this is main") //main 结束了 由main函数启动的goroutine 也都结束了所以要一个等待goroutine
	time.Sleep(time.Second)     //等待goroutine 结束  但是这个是最笨的办法  因为time.Second 可能多可能少不够
}
