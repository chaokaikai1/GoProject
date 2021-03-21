package main

import "fmt"

//有返回值的  声明返回值变量 加括号（ret  int）
//函数体里边可以直接用变量  最后直接return
//return  可以返回值加变量名返回 也可以不加
func sum(a int, b int) (ret int) {
	ret = a + b
	return
}

//返回值的变量也可以不用声明
//参数同类型的 可以用逗号隔开一起声明

func f1(age, high int, name, tel string) string {
	ret := fmt.Sprintf("age=%d,high=%d,name=%s,tel=%s", age, high, name, tel)
	return ret
}

//可变参数  ...
func f2(x, y int, z ...string) {

}

//多返回值 要么都命名  要么都不命名

func f3() (name string, age int) {
	name = "aa"
	age = 10
	return
}

func f4() (string, int) {
	name := "aa"
	age := 10
	return name, age
}

func main() {
	ret := f1(1, 175, "aa", "15011540049")
	fmt.Println(ret)
	fmt.Println(f4())
}
