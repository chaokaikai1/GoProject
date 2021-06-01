package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//值接收者
func (node treeNode) print() {
	node.value = 3
	fmt.Println(node.value)
}

//指针接受者
func (node *treeNode) print2() {
	node.value = 4
	fmt.Println(node.value)
}

func main() {
	//node:=treeNode{
	//	value:1,
	//	left:&treeNode{value:2},
	//}
	//node.print()
	//fmt.Println(node.value)
	//node.print2()
	//fmt.Println(node.value)
	pNode := &treeNode{value: 5}
	pNode.print()
	pNode.print2()
}
