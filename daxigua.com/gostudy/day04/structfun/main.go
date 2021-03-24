package main
import "fmt"

type person struct{
	name string
	age int
}
//结构体的构造函数
//尽量返回值用指针类型 因为 如果返回结构体 值类型 是复制粘贴 会耗内存
//用指针类型 是引用类型
func newPerson(name string, age int) *person{
	return &person{
		name:name,
		age:age,
	}
}
func main() {
	p1:=newPerson("aa",11)
	p2:=newPerson("bb",22)
	
	fmt.Printf("%T\n",p1)
	fmt.Printf("%T\n",p2)
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println(p1.name)
	//fmt.Println(p1.name)
	fmt.Println(p2)
}