package main

import "fmt"

type BookInfo struct {
	title  string
	author string
	price  float32
	id     int
}

func main() {
	bookinfo := BookInfo{title: "钢铁", author: "保尔", price: 32.5, id: 1}

	var book1 BookInfo
	book1.author = "aa"
	book1.title = "我是标题"
	//声明结构指针
	ptrBook := &book1
	fmt.Println(ptrBook)
	fmt.Println(book1.title)
	fmt.Println(bookinfo)
	printBook(bookinfo)
}

func printBook(book BookInfo) {
	fmt.Println(book.author)
}
