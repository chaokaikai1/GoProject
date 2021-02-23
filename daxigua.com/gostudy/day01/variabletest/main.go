package main

import "fmt"

//因式分解关键字的写法 只能用于声明全局变量
var (
	x    int = 1
	y, z string
	q    bool
)

func main() {
	var a, b int = 1, 2
	var name string = "ckk"
	var tel = "1501154049"
	var fmtMsg = "my name is %s and tel is %s and age is %d"
	var result = fmt.Sprintf(fmtMsg, name, tel, a)
	var age int
	newName := "new name"
	var r, w, e = 1, 2, 3
	fmt.Println(r + w + e)
	fmt.Println(newName)
	fmt.Printf("my age is %d", age)
	fmt.Print(age)
	fmt.Println(result)
	fmt.Println(a + b)
	fmt.Println(name)
	fmt.Println("aa")
	fmt.Println("****************************************")

	//多变量声明   := 不带声明格式的只能在函数体内出现
	addr, phone, email := "aaaa", "15011540049", "331079902@qq.com"
	fmtMsg2 := "addr is %s phone is %s email is %s"
	fmt.Printf(fmtMsg2, addr, phone, email)
	fmt.Println()
	fmt.Println(x)
	y = "gy"
	z = "gz"
	q = true
	fmt.Println(y, z, q)

}
