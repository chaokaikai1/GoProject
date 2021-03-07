package main

import "fmt"

func main() {
	//切片和数组声明 基本上一样 无非是 不用声明长度
	//切片是引用类型 数组是值类型
	var s1 []int
	s1 = []int{1, 2, 3}
	fmt.Println(s1)

	//数组得到切片
	numsArry := [...]int{1, 3, 5, 7, 9, 11}
	s2 := numsArry[:2]  //等同于 numsArry[0:2] 从0个开始 切到索引1 左闭 右开 {1,3}
	s3 := numsArry[:]   //等同于从0 切到末尾  len(numsArry)
	s4 := numsArry[2:4] //等同于从索引2开始 切到索引4-1
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	//切片的len 就是len   容量cap是指底层数组的容量(前提是切片从0开始切)
	//len 2  cap= 6
	fmt.Printf("len(s2):%d  ,cap(s2):%d", len(s2), cap(s2))
	fmt.Println()
	//切片的len 就是len   假如切片是从中间切的 那么容量就是底层数组 中间索引到最后的长度
	//len=2  cap=4
	fmt.Printf("len(s4):%d  ,cap(s4):%d", len(s4), cap(s4))

	//切片再切片
	s5 := s4[1:] //  s4 [5,7]    s5  [7]
	fmt.Println()
	fmt.Printf("len(s5):%d  ,cap(s5):%d", len(s5), cap(s5))
}
