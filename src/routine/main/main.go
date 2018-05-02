package main

import (
	"fmt"
	"../eventcenter"
)

func console(interface{}) {
	fmt.Println("console by event")
}

func main() {
	
	event := eventcenter.EventCenter{List: make(map[string]func(interface{}))}
	event.Add("console", console)
	event.Fire("console")
}
