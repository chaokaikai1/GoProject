package main

import (
	"fmt"
)

func main() {
	nameArry := []string{"aa", "bb"}
	var ageArry []int = []int{1, 2}
	var balance = [...]float32{1.0, 2.0}
	balance2 := [2]string{"ccc", "dddd"}
	fmt.Println(ageArry[0])
	fmt.Println(nameArry[0])
	fmt.Println(balance[0])
	fmt.Println(balance2[0])
}
