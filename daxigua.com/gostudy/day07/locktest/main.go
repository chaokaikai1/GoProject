package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var lock sync.Mutex
var x int

func add() {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		lock.Lock() //加锁  互斥锁  同一时间只有一个goroutine 进入临界区其他的gorountine 等待
		x = x + 1
		lock.Unlock() //解锁 //lock 释放之后 其他的goroutine 进入临界区
	}

}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
