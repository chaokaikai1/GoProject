package main

import "fmt"

func maptest1() {
	var m1 map[string]string
	m1 = make(map[string]string, 1)
	m1["name"] = "aa"
	m1["age"] = "22"
	fmt.Println(m1)
}

func maptest2() {
	m1 := make(map[string]string, 1)
	m1["name"] = "ckk"
	m1["addr"] = "beijing"
	m1["tel"] = "15011540049"
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	for _, v := range m1 {
		fmt.Println(v)
	}
}

func maptest3() {
	m1 := make(map[string]string)
	m1["name"] = "ckk"
	m1["addr"] = "beijing"
	m1["tel"] = "15011540049"
	v, ok := m1["hobby"]
	fmt.Println(v, ok) //v 是该索引返回的值，ok是返回true  or false  是否存在该索引

}

func deltest() {
	m1 := make(map[string]string)
	m1["name"] = "ckk"
	m1["addr"] = "beijing"
	delete(m1, "name")
	fmt.Printf("v: %v ,len:%d ", m1, len(m1))
}
func main() {
	//maptest1()
	//maptest2()
	//maptest3()
	deltest()
}
