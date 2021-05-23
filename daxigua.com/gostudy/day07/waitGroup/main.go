package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() //等同于 wg.Add(-1)  线程数-1
	fmt.Printf("hello %v \n", i)
}

func main() {
	wg.Add(10) //说明开启10个线程 gorountie
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	fmt.Println(" over mian")
	wg.Wait() //等待10个线程全部结束 再往下执行
	fmt.Println("啦啦啦")
}
