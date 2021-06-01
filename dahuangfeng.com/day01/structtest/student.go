package main

import "fmt"

type treeNode struct {
	id          int
	name        string
	left, right *treeNode
}

func test1() {
	//也可以这样声明
	//root2:=treeNode{}
	//fmt.Println(root2)

	var root treeNode
	root = treeNode{
		id:   1,
		name: "aa",
		left: &treeNode{},
	}
	//如果想不写key 那么得需要把所有的值给赋值上去
	root.right = &treeNode{2, "bb", nil, nil}
	root.right.left = &root
	fmt.Println(root)
}

func createNode(id int) *treeNode {
	return &treeNode{id: id}
}

func test2() {
	root := createNode(1)
	fmt.Println(root)
}
func test3() {
	nodes := []treeNode{
		{id: 1, name: "aa"},
		{id: 2, name: "bb"},
	}
	fmt.Println(nodes)
}

func main() {
	//test1()
	test3()
}
