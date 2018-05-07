package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
	Age  int8
}

func main() {
	time := time.Unix()
	fmt.Println(time)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
