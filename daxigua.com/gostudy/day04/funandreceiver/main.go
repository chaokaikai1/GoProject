package main

import "fmt"
//给自定义类型加方法

type myInt int

type dog struct{
	name string 
	age int
}

type person struct{
	name string 
	age int
}
//在方法前边加括号 然后声明类型以及类型变量 表明该方法只能被该类型调用
//通常变量为该类型首字母的小写
//值接收者
func (d dog)Eat()  {
	fmt.Printf("%s is eat \n",d.name)
}
//值接受者 dog的年龄不变
func  (d dog) guonian(){
	d.age++
}

func (m myInt) hello(){
	fmt.Println("hello")
}

//指针接收者
//引用地址  会将原始值改变
func (p *person) guonian(){
	p.age++
}

func main() {
	d1:=dog{
		name:"aa",
		age:1,
	}
	d1.Eat()
	d1.guonian() //1
	fmt.Println(d1.age)
	p1:=person{
		name:"bb",
		age:1,
	}
	p1.guonian()
	fmt.Println(p1) //{bb 2}

	m:=myInt(100)
	fmt.Println(m)
	m.hello()
}