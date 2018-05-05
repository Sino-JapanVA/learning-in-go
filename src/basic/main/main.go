package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	p, err := ReadFrom(strings.NewReader("Treasure"), 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}
