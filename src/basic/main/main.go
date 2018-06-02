package main

import (
	"fmt"
)

// User 用户
type User struct {
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	Favorite []Book
}

// Book 书籍
type Book struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {
	// jsonBytes, err := ioutil.ReadFile("../test.json")
	// checkErr(err)
	// // jsonText := string(jsonBytes)
	// var user User
	// err = jsoniter.Unmarshal(jsonBytes, &user)
	// checkErr(err)
	// fmt.Println(user.Favorite)
	// txt, err := jsoniter.Marshal(user)
	// fmt.Println(string(txt))
	// data := []byte(string(txt))
	// fmt.Println(data)
	back()
	fmt.Println("after func")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func back() bool {
	return true
}
