package main

import (
	"fmt"
	"strings"
)

func main() {

	//字符串转义
	path := "D:\\Program Files\\Notepad++\\localization"
	//反引号 原样输出
	path2 := `D:\Program Files\Notepad++\localization`
	fmt.Println(path)
	fmt.Println(path2)
	//字符串常用操作
	//len 长度
	fmt.Println(len(path))
	//拼接 用+  或者 fmt.Sprintf 返回字符串
	name := "ckk"
	age := "30"
	result1 := name + age
	result2 := fmt.Sprintf("%s %s", name, age)
	fmt.Println(result2)
	fmt.Println(result1)

	//字符串分割
	pathArry := strings.Split(path, "\\")
	fmt.Println(pathArry)
	fmt.Println(pathArry[0])
	fmt.Println(len(pathArry))

	//字符串包含
	fmt.Println(strings.Contains(name, "a"))

	//索引
	index := strings.Index(path, "P")
	fmt.Println(index)

	//前缀 后缀
	b1 := strings.HasPrefix(name, "a")
	b2 := strings.HasSuffix(name, "k")
	fmt.Println(b1)
	fmt.Println(b2)

	//拼接
	resultJoin := strings.Join(pathArry, ",")
	fmt.Println(resultJoin)

}
