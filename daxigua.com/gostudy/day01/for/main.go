package main

import "fmt"

func main() {
	sum := 0
	nameArry := []string{"aa", "bb", "cc"}
	//普通循环
	for i := 0; i < 10; i++ {
		sum += i
	}

	//for 变种1  可以不输入初始的 从上边找
	j := 0
	for ; j < 10; j++ {
		if j == 5 {
			break //跳出循环
		}
		if j == 4 {
			continue
		}
		fmt.Println(j)
	}

	//循环数组
	for key, item := range nameArry {
		fmt.Println(key, item)
	}
	fmt.Println(sum)

	//多层循环 想跳出循环 可以用goto 到标签 跳出

	for i := 0; i < 5; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'B' {
				goto breakLabel
			}
			fmt.Printf("%d %c", i, j)
		}
	}
breakLabel:
}
