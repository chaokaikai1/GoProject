package main

import "fmt"

func main() {
	//make  指定切片类型 长度 以及容量 如果不声明容量 那么默认容量和长度相等
	//5 是len   10 是cap
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v  len(s1)=%d   cap(s1)=%d\n", s1, len(s1), cap(s1))
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
	}
	fmt.Printf("s1=%v  len(s1)=%d   cap(s1)=%d\n", s1, len(s1), cap(s1))

	//数组  数值类型  类比 word文档的副本  源文件改变 副本内容不变
	//切片 引用类型  类比理解  word的快捷方式  源文件改变 快捷方式打开看到的也变化
	//切片指针 指向引用地址
	a1 := [2]int{5, 7}
	a2 := a1
	fmt.Println(a1, a2) //{5,7}  {5,7}
	a1[0] = 10
	fmt.Println(a1, a2) //{10,7}  {5,7}
	s2 := []int{1, 3}
	s3 := s2
	fmt.Println(s2, s3) //{1,3}  {1,3}
	s2[0] = 10
	fmt.Println(s2, s3) //{10,3}{10,3}
}
