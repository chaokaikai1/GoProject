package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "1000"
	//Atoi 字符串转int
	i1, _ := strconv.Atoi(str1)
	fmt.Println(i1)
	//Itoa int 转字符串
	str2 := strconv.Itoa(i1)
	fmt.Println(str2)
	//字符串转float
	f1, _ := strconv.ParseFloat(str1, 64)
	fmt.Println(f1)

	b1str := "true"
	b1, _ := strconv.ParseBool(b1str)
	fmt.Println(b1)
}
