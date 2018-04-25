package main

import "fmt"

func main() {
	var a int32 = 256
	var b int16 = 8
	var c = int(a) + int(b)
	fmt.Println(c)
}
