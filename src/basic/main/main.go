package main

import (
	"basic/List"
	"basic/Stack"
	"fmt"
	"math/rand"
)

func main() {
	funcStack()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func funcList() {
	list := List.NewArrayList(20)
	for i := 0; i < 10; i++ {
		list.Append(rand.Intn(10))
	}
	fmt.Println("ArrayList")
	// list.Print()

	nodeList := List.NewNodeList()
	for i := 0; i <= 10; i++ {
		nodeList.Append(i)
	}
	fmt.Println("NodeList")

	fmt.Println(nodeList.Size())

	// 找到值
	var ele int
	nodeList.GetVal(&ele)
	// nodeList.Print()

	// fmt.Println("Find ele")
	// var ele2 *List.Node
	// 指向指针的指针
	// isFind := nodeList.Find(&ele2, 5)
	// if isFind {
	// 	ele2.Val = 199
	// 	fmt.Println(ele2.Val)
	// } else {
	// 	fmt.Println("can not fin val")
	// }

	nodeList.Remove(5)
	nodeList.Print()

	nodeList.Insert(5)
	nodeList.Insert(5)
	nodeList.Print()
}

func funcStack() {
	stack := Stack.NewStack(10)
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	val1 := stack.GetValue()
	fmt.Println(val1)

	stack.Print()

	var ele int
	stack.Pop(&ele)
	val2 := stack.GetValue()
	fmt.Println(ele, val2)
	stack.Print()
}

func funcSliceList() {
	sliceList := List.NewSliceList()
	sliceList.Append(1)
	sliceList.Append(2)
	sliceList.Append(3)

	sliceList.Insert(4, 2)
	var ele int
	sliceList.Remove(&ele, 0)

	sliceList.Print()
}
