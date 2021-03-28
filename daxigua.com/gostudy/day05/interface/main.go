package main

import "fmt"

//phone接口
//不管什么类型  只要有call（实现接口里所有的方法） 方法 都是继承 phone类型（接口）
type phone interface {
	call()
}

//再声明一个接受接口类型的方法调用接口方法
func phoneCall(p phone) {
	p.call()
}

type nokia struct {
	name string
}

func (n nokia) call() {
	fmt.Printf("%s is call \n", n.name)
}

type motorola struct {
	name string
}

func (m motorola) call() {
	fmt.Printf("%s is call \n", m.name)
}
func main() {
	n1 := nokia{name: "nokiaPhone"}
	m1 := motorola{name: "motoralaPhone"}

	//phone1:=new (nokia)
	//phone1.call()
	//phone2:=new(motorola)
	//phone2.call()
	var p1 phone = n1
	p1.call()
	var p2 phone = m1
	p2.call()
	phoneCall(n1)
	phoneCall(m1)

}
