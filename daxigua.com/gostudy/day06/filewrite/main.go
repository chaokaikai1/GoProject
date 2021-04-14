package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("openfile error ,err:%v", err)
		return
	}
	fileObj.WriteString("啦啦啦啦")
	fileObj.Close()
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("openfile error ,err:%v", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("我是添加的字符串") //写到缓存
	wr.Flush()                 //将缓存的写到文件
}

func writeDemo3() {
	str := "我是ioutil 添加的字符串"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("writefile error ,err:%v", err)
		return
	}
}
func main() {
	writeDemo3()
}
