package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	defer wg.Wait()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			wg.Done()
			fmt.Println(i)
		}(i)
	}
}
