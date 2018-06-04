package Dict

import (
	"fmt"
	"strconv"
)

type Dict struct {
	container map[string]int
}

func NewDict() *Dict {
	return &Dict{make(map[string]int)}
}

func (this *Dict) Clear() {
	for key, _ := range this.container {
		delete(this.container, key)
	}
}

func (this *Dict) Insert(key string, val int) {
	if _, ok := this.container[key]; ok {
		this.container[key] = val
	} else {
		this.container[key] = val
	}
}

func (this *Dict) Remove(key string, ele *int) bool {
	if val, ok := this.container[key]; ok {
		*ele = val
		delete(this.container, key)
		return true
	} else {
		return false
	}
}

func (this *Dict) RemoveAny() {

}

func (this *Dict) Find(key string, ele *int) bool {
	if val, ok := this.container[key]; ok {
		*ele = val
		return true
	} else {
		return false
	}
}

func (this *Dict) Print() {
	for key, val := range this.container {
		fmt.Println(key + ":" + strconv.Itoa(val))
	}
}

func (this *Dict) Size() int {
	return len(this.container)
}
