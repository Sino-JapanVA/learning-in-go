package main

import (
	"fmt"
)

// EventCenter 事件中心
type EventCenter struct {
	list map[string]func(interface{})
}

func (e *EventCenter) add(name string, fun func(interface{})) {
	e.list[name] = fun
}

func (e *EventCenter) fire(name string) {
	fun, err := e.list[name]
	if !err {
		fmt.Println("Error in func", err)
		return
	}
	fun(nil)
}

func console(interface{}) {
	fmt.Println("console by event")
}

func main() {
	event := EventCenter{list: make(map[string]func(interface{}))}
	event.add("console", console)
	event.fire("console")
}
