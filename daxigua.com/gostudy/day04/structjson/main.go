package main

import (
	"encoding/json"
	"fmt"
)

//struct 和json的转换

//如果struct 字段要和json转换那么 字段必须声明为首字母大写 这样相当于public 才可以被外部包json  访问到序列化
//如果该字段返回名回想变成其他的名称 可以用tag  tag在这里的用途就是提供json结构中的别名
type person struct {
	Name string `json:"myname"`
	Age  int    `json:"age"`
}

type student struct {
	Name string
	Age  int
}

func main() {
	p1 := person{
		Name: "aa",
		Age:  11,
	}
	data, _ := json.Marshal(p1)
	fmt.Println(string(data))
	fmt.Printf("%#v \n", string(data)) //%#v %v相应值默认格式   %#v 相应值 go语言中的格式

	s1 := `{"Name":"bb","Age":22}` //反斜线原样输出

	var p2 person
	json.Unmarshal([]byte(s1), &p2)
	fmt.Printf("%#v", p2)

	st1 := student{
		Name: "啦啦啦",
		Age:  22,
	}
	data2, _ := json.Marshal(st1)
	fmt.Println(string(data2))

	str1 := `{"Name":"哈哈哈哈哈","Age":33}`
	st2 := student{}
	json.Unmarshal([]byte(str1), &st2)
	fmt.Printf("st2 %v", st2)

}
