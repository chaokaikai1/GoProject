package main

import "fmt"

func test1() {
	slice1 := []int{1, 3, 5, 7, 9}
	slice2 := make([]int, 3)    //len 和cap都是3  输出 [0,0,0]
	slice3 := make([]int, 3, 4) //len3 cap 4
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
func appendtest() {
	slice1 := []int{1, 3, 5, 7, 9}
	s1 := append(slice1, 11)
	fmt.Println(s1)
}

//如果加入的两个数组切片不一样大，就会按其中较小的那个数组切片的元素个数进行复制。
func copytest() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println(slice2)  //此时的slice2 shi [1,2,3]
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)  //[5,4,3,4,5]
}

func deltest() {
	s1 := []int{1, 3, 5, 7, 9}
	s2 := append(s1[1:2], s1[3:5]...) //...如果append 后边追加的是切片，需要加...可变长度
	fmt.Printf("v %v,len %d,cap %d\n", s2, len(s2), cap(s2))
}

func main() {
	//test1()
	//appendtest()
	//copytest()
	deltest()
}
