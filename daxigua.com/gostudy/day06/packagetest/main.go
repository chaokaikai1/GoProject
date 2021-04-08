package main

import (
	"fmt"
	cals "gostudy/day06/01pack"
	"gostudy/day06/testpack"
)

func main() {
	ret := testpack.Add(1, 2)
	ret2 := cals.Add(1, 2)

	fmt.Println(ret)
	fmt.Println(ret2)
}
