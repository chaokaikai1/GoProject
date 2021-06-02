package main

import (
	real2 "GoProject/dahuangfeng.com/day02/interfacetest/retriever/real"
	"fmt"
)

type IRetriever interface {
	Get(url string) string
}

func download(r IRetriever) string {
	return r.Get("http://www.imooc.com")
}
func main() {
	//url:="http://www.imooc.com"
	//re:=mock.Retriever{Contents: "this is a fake imooc.com"}
	re := real2.Retriever{}
	ret := download(re)
	fmt.Println(ret)
}
