package main

import (
	"fmt"
	"io"
)

func main() {
	str := "Sunshine夏青"
	fmt.Println(len(str))
	fmt.Println(len([]byte(str)))
	slice := []rune(str)
	slice[len(slice)-1] = '日'
	fmt.Println(len([]rune(str)))
	fmt.Println(slice)
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}
