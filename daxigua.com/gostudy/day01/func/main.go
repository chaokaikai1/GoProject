package main

import "fmt"

func main() {
	name := "ckk"
	age := 11
	a, b := 1, 2
	nameDes := fmtNameDes(name, age)
	addResult := add(&a, &b)
	fmt.Println(nameDes)
	fmt.Println(addResult)
}

//格式化名称描述
func fmtNameDes(name string, age int) string {
	fmtMsg := "my name is %s and age is %d"
	result := fmt.Sprintf(fmtMsg, name, age)
	return result
}

//加法 引用传递
func add(a *int, b *int) int {
	*a = 3
	*b = 4
	return *a + *b
}
