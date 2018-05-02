package eventcenter

import "fmt"

// EventCenter 事件中心
type EventCenter struct {
	List map[string]func(interface{})
}

// Add 注册事件
func (e *EventCenter) Add(name string, fun func(interface{})) {
	e.List[name] = fun
}

// Fire 触发事件
func (e *EventCenter) Fire(name string) {
	fun, err := e.List[name]
	if !err {
		fmt.Println("Error in func", err)
		return
	}
	fun(nil)
}
