package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	array2 := [...]int{1, 2}
	slice := array[0:3:3]
	fmt.Println(slice)
	slice = append(slice, 11)
	slice[0] = 0
	fmt.Println(slice, cap(slice))
	fmt.Println(array)
}
