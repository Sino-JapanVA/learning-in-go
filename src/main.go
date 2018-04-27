package main

import (
	"fmt"
)

// Demo
type Demo struct {
	x int
	y int
	z string
}

func main() {

	demo := &Demo{x: 1, y: 2, z: "Hello Go"}
	fmt.Printf("%d\t%d\t%s\n", demo.x, demo.y, demo.z)
	change(demo)
	fmt.Printf("%d\t%d\t%s\n", demo.x, demo.y, demo.z)
}

func change(demo *Demo) {
	demo.x = 2
	demo.y = 3
	demo.z = "Hello World"
}
