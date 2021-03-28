package main

import (
	"encoding/json"
	"fmt"
)

type dog struct {
	name string
	age  int
}

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (d dog) eat() {
	fmt.Printf("dog name is  %s  ,say wangwang \n", d.name)
}
func main() {
	d1 := dog{
		name: "aa",
	}
	d1.eat()

	arry1 := [...]string{"aaa", "bbb", "ccc"}
	for _, v := range arry1 {
		fmt.Println(v)
	}
	s1 := student{
		Name: "aa",
		Age:  11,
	}
	data, _ := json.Marshal(s1)
	fmt.Println(string(data))
	fmt.Printf("%v\n", string(data))
	fmt.Printf("%#v\n", string(data))

	var stu student
	s2 := `{"name":"bb","age":22}`
	json.Unmarshal([]byte(s2), &stu)
	fmt.Println(stu)
}
