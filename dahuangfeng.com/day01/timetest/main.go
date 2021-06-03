package main

import (
	"fmt"
	"time"
)

func main() {
	t1:= time.Now()
	fmt.Println(t1)
	t2:=time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Println(t2)
    fmt.Println(time.Now().Format("2006-01-02"))

}
