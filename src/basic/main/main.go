package main

import (
	"strings"
	"basic/List"
	"basic/Stack"
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

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

// 逆波兰算法相关
func postFix() {
	nums := []string{"9", "3", "1", "-", "3", "*", "+", "10", "2", "/", "+"}
	stack := Stack.NewStack(20)

	for _, val := range nums {
		if isOperation(val) {
			var prev, next int
			stack.Pop(&next)
			stack.Pop(&prev)
			result := operation(prev, next, val)
			stack.Push(result)
		} else {
			num, _ := strconv.Atoi(val)
			stack.Push(num)
		}
	}
	fmt.Println(stack.GetValue())
}
func pre2Post() *[]string {
	slice := strings.Split("")
	stack := Stack.NewStack(20)
	var slice2 []string
	for _, val := range slice {
		if isOperation(val) {
			if (val != ")") {
				now := stack.GetValue()
				if now == "+" || now == "-" {
					stack.Push(val)
				} else {
					var popEle int
					stack.Pop(&popEle)
					slice2 = append(slice2, popEle)
				}

			} else {
				for ele := stack.Pop(&ele); ele != "(" {
					slice2 = append(slice2, ele)
				}
			}
		} else {
			slice2 = append(slice2, val)
		}
	}
}

func operation(a, b int, opt string) int {
	switch opt {
	case "-":
		return a - b
	case "+":
		return a + b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("wrong operation")

	}
}

func isOperation(str string) bool {
	slice := []string{"+", "-", "*", "/", "(", ")"}
	for _, val := range slice {
		if val == str {
			return true
		}
	}
	return false
}
