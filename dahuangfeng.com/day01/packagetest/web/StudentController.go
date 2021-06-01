package main

import (
	"GoProject/dahuangfeng.com/day01/packagetest/entity"
	"fmt"
)

func main() {
	stu := entity.Student{Id: 1, Age: 10, Name: "aa"}
	stu.PrintAge()
	age := stu.SumAge()
	fmt.Println(age)
}
