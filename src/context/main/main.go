package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/context"
)

func main() {
	http.Handle("/", http.HandlerFunc(myHandler))
	http.ListenAndServe(":1234", nil)
}

func myHandler(rw http.ResponseWriter, r *http.Request) {
	context.Set(r, "user", "Treasure")
	context.Set(r, "age", 22)

	doHandler(rw, r)
}

func doHandler(rw http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user").(string)
	age := context.Get(r, "age").(int)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("The user is " + user + ", age is " + strconv.Itoa(age)))
}
