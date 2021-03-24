package main

import "fmt"
type person struct{
	name string
	age int
}
func main() {
	//结构体指针声明
	//new  声明值类型的 指针 
	//make  创建 切片slice  map  chanel
	var p=new(person)
	p.name="aa"
	p.age=11
	fmt.Printf("%T\n",p)
	fmt.Println(p)

	//声明指针方式2 key  value

	var p2=&person{
		name:"bb",
		age:22,
	}
	
	fmt.Printf("%T\n",p2)
	fmt.Println(p2)

	p3:=&person{
		name:"cc",
		age:33,
	}
	fmt.Println(p3)

}