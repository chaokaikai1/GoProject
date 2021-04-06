package main

import "fmt"

func show(a interface{}) {
	fmt.Printf("type is %T , value is %#v \n", a, a)
}

func assign(a interface{}) {

	switch v := a.(type) {
	case string:
		fmt.Println("type is 字符串 ", v)
	case int:
		fmt.Println("type is int", v)
	case bool:
		fmt.Println("type is bool", v)
	}
}
func main() {
	//空接口即 没有方法的接口
	//空接口用法1  用于声明map的值
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 6)
	m1["name"] = "aa"
	m1["age"] = 11
	m1["hobby"] = [...]string{"唱歌", "跳舞"}
	fmt.Println(m1)
	//空接口用法2  方法参数的声明
	show(false)
	show(1)
	show("aa")

	//空接口 用法3 类型断言
	assign("bbb")
}
