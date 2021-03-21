package main

import "fmt"

//defer 多用于函数结束之前 释放资源（数据库连接  文件关闭 socket连接等）
func defer_demo() {
	fmt.Println("start")
	defer fmt.Println("aaa") //defer 后边的语句会在函数执行完之前执行
	defer fmt.Println("bbb") //多个deffer  采用后进先出的方式
	defer fmt.Println("ccc") //start  end    ccc  bbb aaa
	fmt.Println("end")
}

func main() {
	defer_demo()
}
