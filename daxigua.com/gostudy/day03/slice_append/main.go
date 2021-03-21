package main

import "fmt"

func main() {
	//切片append 的策略  如果新申请的容量大于旧容量2倍 则用新申请的容量
	//如果旧切片长度小于1024 则是旧容量的2倍
	//如果旧切片长度大于等于1024 则新申请的容量 每次增加1/4  直到最终容量大于等于新申请的容量
	//如果最终容量计算溢出 则最终容量就是新申请的容量
	s1 := []string{"北京", "上海"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "广州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := append(s1, "深圳")
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))
	s3 := []string{"苏州", "西安"}
	s1 = append(s1, s3...) //...表示切片拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

}
