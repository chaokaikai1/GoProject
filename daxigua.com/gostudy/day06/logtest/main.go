package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file error ,err:%v \n", err)
		return
	}
	log.SetOutput(fileObj)
	log.Println("我是记录的日志啊")
}
