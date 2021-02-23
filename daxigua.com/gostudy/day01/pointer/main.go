package main

import "fmt"

func main() {
	// var a int = 10
	// name := "ckk"
	// fmt.Println(&a)
	// fmt.Println(&name)

	a := 1
	name := "aa"
	ip := &a
	var ip2 *int = &a
	var ip3 *string = &name
	var ptr *int
	bResult := ptr == nil
	fmt.Println(ip)
	fmt.Println(ip2)
	fmt.Println(ip3)
	fmt.Println(ptr)
	fmt.Println(bResult)
}
