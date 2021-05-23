package main

import (
	"fmt"
	"time"
)

//声明一个workerpool
//单向通道  <-chan  只读从通道里读取     chan<-  只写  往通道里写
func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Printf("worker %d start job :%d\n", id, i)
		time.Sleep(time.Second)
		fmt.Printf("worker %d end job :%d\n", id, i)
		results <- i * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启3个gorountine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}
	//5个任务
	for j := 0; j < 5; j++ {
		jobs <- j //将j  放进chan里
	}
	close(jobs)
	//输出结果
	for a := 0; a < 5; a++ {
		<-results
	}
}
