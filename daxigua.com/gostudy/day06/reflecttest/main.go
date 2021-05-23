package main

import (
	"fmt"
	"reflect"
)

func rtypeof(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("x type is %v\n", t)
	fmt.Printf("type  name is %v ,type kind is %v", t.Name(), t.Kind())
}
func rvalueof(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("x value is %v", v)

}

//反射根据传递过来的值 赋值
func rSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	//反射中使用elem 获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 || v.Elem().Kind() == reflect.Int32 {
		v.Elem().SetInt(200)
	}

}

type student struct{}

func main() {
	s1 := student{}
	rtypeof(s1)
	var v int32 = 2
	rSetValue(&v)
	fmt.Println(v)
}
