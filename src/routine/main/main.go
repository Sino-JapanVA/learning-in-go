package main

import (
	"fmt"
	event "routine/eventcenter"
)

func console(interface{}) {
	fmt.Println("console by event")
}

func main() {

	event := event.EventCenter{List: make(map[string]func(interface{}))}
	event.Add("console", console)
	event.Fire("console")
}
