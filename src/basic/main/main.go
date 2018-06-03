package main

import (
	"basic/List"
	"fmt"
	"math/rand"
)

func main() {

	list := List.NewArrayList(20)
	for i := 0; i < 10; i++ {
		list.Append(rand.Intn(10))
	}
	fmt.Println("ArrayList")
	// list.Print()

	nodeList := List.NewNodeList()
	for i := 0; i < 10; i++ {
		nodeList.Append(rand.Intn(10))
	}
	fmt.Println("NodeList")

	fmt.Println(nodeList.Size())
	var ele int

	nodeList.GetVal(&ele)
	nodeList.Print()

	fmt.Println("Find ele")
	var ele2 *List.Node
	// 指向指针的指针
	isFind := nodeList.Find(&ele2, 5)
	if isFind {
		ele2.Val = 199
		fmt.Println(ele2.Val)
	} else {
		fmt.Println("can not fin val")
	}
	nodeList.Print()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
