package main

import (
	"fmt"
	"io/ioutil"

	"github.com/json-iterator/go"
)

type User struct {
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	Favorite []Book
}

type Book struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {
	jsonBytes, err := ioutil.ReadFile("../test.json")
	checkErr(err)
	// jsonText := string(jsonBytes)
	var user User
	err = jsoniter.Unmarshal(jsonBytes, &user)
	checkErr(err)
	fmt.Println(user.Favorite)
	txt, err := jsoniter.Marshal(user)
	fmt.Println(string(txt))
	data := []byte(string(txt))
	fmt.Println(data)
	fmt.Println(time)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
