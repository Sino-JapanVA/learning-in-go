package main

import (
	"container/list"
	"fmt"
)

type User struct {
	Name string
	Age  int8
}

func main() {
	list := list.New()
	list.PushBack(User{"Treasure", 20})
	ele := list.Front()
	fmt.Println(ele.Value)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
