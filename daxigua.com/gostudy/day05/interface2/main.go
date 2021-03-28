package main

import "fmt"

type animal interface {
	say()
	eat(string)
	move() string
}

type cat struct {
	name string
}

type dog struct {
	name string
}

//指针接收者  一般用的多  在调用的时候 记得 声明接口指向指针
func (d *dog) say() {

}
func (d *dog) eat(food string) {
	fmt.Printf("%s eat %s", d.name, food)
}
func (d *dog) move() string {
	return "aaa"
}

//值接受者   既可以声明接口指向struct  也可以指向指针
func (c cat) say() {
	fmt.Println("cat say miao miao")
}
func (c cat) eat(food string) {
	fmt.Printf("%s is eat %s \n", c.name, food)
}
func (c cat) move() string {
	return "我是返回值"
}
func main() {
	c1 := cat{
		name: "豆丁",
	}
	var ani animal
	var a1 animal
	ani = c1
	a1 = &c1 //值接受者 既可以指向指针也可以指向struct 对象
	a1.eat("啦啦啦")
	ani.eat("小黄鱼")
	fmt.Println(ani.move())
	ani.say()
	//接口保存 分两部分  一部分 是动态类型  一部分是动态的值    只有这样接口才可以接收各种类型
	fmt.Printf("%T \n", ani) //main.cat

	d1 := dog{name: "tom"}
	var a2 animal
	a2 = &d1 //指针接受者只能指向指针
	a2.eat("屎")

}
