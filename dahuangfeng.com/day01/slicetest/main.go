package main

import "fmt"

func slice1test() {
	//数组
	arry1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//切片就是把[...]换成[]int 或者在数组上边  [:] 截取 [2:5] 左包含 右不包含
	slice1 := arry1[2:5]
	sliceend := arry1[2:] //从索引位置截取到最后
	slice3 := arry1[:3]   //从0到2   左包含右不包含
	fmt.Println(slice1)
	fmt.Println(sliceend)
	fmt.Println(slice3)

}

func slice2test() {
	//slice1:=[]int {0,1,2,3}
	//slice2:=slice1[2:5]
	//fmt.Println(slice2)// slice bounds out of range [:5] with capacity 4
	//

	arry1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice1 := arry1[2:6]
	slice2 := slice1[3:5]
	fmt.Printf("slice1  len:%d ,cap :%d, v:%v\n", len(slice1), cap(slice1), slice1) //slice1  len:4 ,cap :6, v:[2 3 4 5]
	fmt.Println(arry1)                                                              //[0,1,2,3,4,5,6,7]
	//slice1:2,3,4,5        slice2[3:5]  底层数组  2,3,4,5,6,7  所以  [3,5] 就是[5,6]
	fmt.Printf("slice2  len:%d ,cap :%d, v:%v\n", len(slice2), cap(slice2), slice2) //slice2  len:2 ,cap :3, v:[5 6]

}

func main() {
	//slice1test()
	slice2test()
}
