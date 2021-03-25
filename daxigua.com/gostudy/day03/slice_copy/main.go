package main

import "fmt"

func main() {
	s1 := []int{1, 3, 5}
	s2 := s1
	//var s3 []int //未初始化的 切片是nil 没法copy
	var s3 = make([]int, 3)
	fmt.Println(s1, s2, s3) //[1 3 5] [1 3 5] [0 0 0]
	copy(s3, s1)
	s1[0] = 100
	fmt.Println(s1, s2, s3) //[100 3 5] [100 3 5] [1 3 5]  copy的 不会更改

	//切片的删除 切片没有删除 可以用切片的分割 在append
	//删除索引为1的元素
	s4 := []int{1, 3, 5, 7, 9}
	s4 = append(s4[:1], s4[2:]...)
	fmt.Printf("s4=%v  len(s4)=%d cap(s4)=%d", s4, len(s4), cap(s4))
	fmt.Println("--------------------")

	//1 切片不保存具体的值
	//2 切片对应一个底层数组
	//3 底层数组都是占用一块连续的内存
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s5 := a1[:]
	fmt.Println(a1, s5) //[1, 3, 5, 7, 9, 11, 13, 15, 17] [1, 3, 5, 7, 9, 11, 13, 15, 17]
	s5 = append(s5[:1], s5[4:6]...)
	//切片改变了底层数组 吧索引为1的3删除 但是切片容量还是那么大 只是吧后增加的数挨个 从索引为1的放进去17 最后一个元素还是17 没动
	fmt.Println(a1, s5) //[1,  5, 7, 9, 11, 13, 15, 17,17] [1, 5, 7, 9, 11, 13, 15, 17]
}