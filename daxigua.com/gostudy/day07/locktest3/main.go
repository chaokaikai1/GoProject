package main

import (
	"fmt"
	"sync"
	"time"
)

var rwLock sync.RWMutex //读写锁
var wg sync.WaitGroup
var x int

func read() {
	defer wg.Done()
	rwLock.RLock()
	println(x)
	rwLock.RUnlock()

}
func write() {
	defer wg.Done()
	rwLock.Lock()
	x = x + 1
	rwLock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	time.Sleep(time.Second)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))

}
