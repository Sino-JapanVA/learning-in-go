package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{"alice": 31, "charlie": 34}
	var p *map[string]int = &ages
	ages["jack"] = 40
	for name, age := range ages {
		fmt.Printf("%s\t %d\n", name, age)
	}
	fmt.Println("\n", *p)
}
