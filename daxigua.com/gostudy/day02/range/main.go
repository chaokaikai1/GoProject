package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 8}
	for key, num := range nums {
		fmt.Println(key, num)
	}
}
