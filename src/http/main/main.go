package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/astaxie/beego/httplib"
)

func main() {
	req := httplib.Get("http://www.baidu.com")

	str, err := req.String()
	checkErr(err)
	fmt.Printf(str)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("-------------------------Request Start----------------------------------")
	fmt.Println("path", r.URL.Path)
	fmt.Println("schema", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
	fmt.Println(r.URL)
	fmt.Println("-------------------------Request End------------------------------------")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Req method: ", r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
