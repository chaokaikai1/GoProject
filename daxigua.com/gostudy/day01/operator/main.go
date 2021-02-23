package main

import "fmt"

func main() {
	name := "ckk"
	a, b := 1, 2
	result := a + b
	isMan := false
	var str string
	//if
	if a > b {
		isMan = true
	} else {
		isMan = false
	}

	//if else if
	if a == 1 {
		str = "1"
	} else if a == 2 {
		str = "2"
	} else {
		str = "default"
	}

	//switch case

	switch a {
	case 1:
		str = "1"
	case 2:
		str = "2"
	case 3:
		str = "3"
	default:
		str = "0"
	}
	fmt.Println(name)
	fmt.Println(result)
	fmt.Println(isMan)
	fmt.Println(str)
}
