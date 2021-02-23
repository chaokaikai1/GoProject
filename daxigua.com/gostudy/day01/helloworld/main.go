package main

import "fmt"

func main() {
	var age = 1
	fmt.Println(age)
	var code = 2
	var name = "aa"
	var url = "http://api.com?code=%d&name=%s"
	var result = fmt.Sprintf(url, code, name)
	fmt.Println(result)
	fmt.Println("hello world")
}
