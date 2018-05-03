package main

import (
	"fmt"
)

type User struct {
	Name string
	age  int8
}

func (this User) setAge(age int8) {
	this.age = age
}

func (this *User) setPointerAge(age int8) {
	this.age = age
}

func main() {
	user := User{Name: "Treasure", age: 22}
	user.setAge(23)
	fmt.Println(user) // {Treasure 22}
	user.setPointerAge(23)
	fmt.Println(user) // {Treasure 23}
}
