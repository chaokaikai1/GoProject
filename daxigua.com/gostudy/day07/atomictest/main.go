package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var lock sync.Mutex
var x int64

func lockAdd() {
	lock.Lock()
	x = x + 1
	lock.Unlock()
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&x, 1) //等价于 锁 和x++
	wg.Done()
}

func main() {
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		// go lockAdd()
		go atomicAdd()
	}
	wg.Wait()
	fmt.Println(x)
}
