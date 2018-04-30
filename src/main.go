package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func main() {
	age := 22
	lisa := user{
		name: "Lisa",
		age:  age}
	fmt.Println(lisa)
}
