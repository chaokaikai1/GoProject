package main

import (
	"fmt"
	"runtime"
	"sync"
)

func a() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("A i%v \n", i)
	}
}

func b() {
	defer wg.Done() //wg.Add(-1)
	for i := 0; i < 100; i++ {
		fmt.Printf("B i%v \n", i)
	}
}

var wg sync.WaitGroup

func main() {
	defer wg.Wait()       //等待所有线程结束
	runtime.GOMAXPROCS(2) //开启cpu核心数  2核
	wg.Add(2)             //开启两个线程
	go a()
	go b()
	fmt.Println("over main")
}
