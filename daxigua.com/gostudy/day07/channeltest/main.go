package main

import (
	"fmt"
	"sync"
)

var c chan int
var wg sync.WaitGroup

func noBufChan() {
	c = make(chan int) //声明 不带缓冲区初始化
	wg.Add(1)
	go func() {
		defer wg.Done() //相当于 wg.Add(-1)
		x := <-c
		fmt.Printf("从c chnanel中取的值 %v \n", x)
		close(c)
	}()
	c <- 10
	fmt.Printf("将10 放到c channel 中 \n")
	fmt.Println("main is over")
	wg.Wait()
}

func bufChan() {
	c = make(chan int, 10) //声明 带缓冲区初始化
	wg.Add(1)
	go func() {
		defer wg.Done() //相当于 wg.Add(-1)
		x := <-c
		fmt.Printf("从c chnanel中取的值 %v \n", x)
		close(c)
	}()
	c <- 10
	fmt.Printf("将10 放到c channel 中 \n")
	fmt.Println("main is over")
	wg.Wait()
}

func main() {
	//从c chnanel中取的值 10
	//将10 放到c channel 中
	//main is over
	//noBufChan()

	bufChan()

}
