package main

import "fmt"

func main() {
	sum := 0
	nameArry := []string{"aa", "bb", "cc"}
	//普通循环
	for i := 0; i < 10; i++ {
		sum += i
	}
	//循环数组
	for key, item := range nameArry {
		fmt.Println(key, item)
	}
	fmt.Println(sum)
}
