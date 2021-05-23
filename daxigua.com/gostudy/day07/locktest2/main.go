package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex //互斥锁
var wg sync.WaitGroup
var x int

func read() {
	defer wg.Done()
	lock.Lock()
	println(x)
	lock.Unlock()

}
func write() {
	defer wg.Done()
	lock.Lock()
	x = x + 1
	lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))

}
