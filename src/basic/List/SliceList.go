package List

import (
	"fmt"
	"strconv"
)

type SliceList struct {
	slice []int
}

func NewSliceList() *SliceList {
	list := []int{}
	return &SliceList{list}
}

func (this *SliceList) Clear() {
	this.slice = []int{}
}

func (this *SliceList) Append(ele int) {
	this.slice = append(this.slice, ele)
}

func (this *SliceList) Insert(ele int, index int) {
	rest := append([]int{}, this.slice[index:]...)
	this.slice = append(this.slice[:index], ele)
	this.slice = append(this.slice, rest...)
}

func (this *SliceList) Remove(ele *int, index int) bool {
	if index > len(this.slice) {
		return false
	} else {
		*ele = this.slice[index]
		this.slice = append(this.slice[:index], this.slice[index+1:]...)
		return true
	}
}

func (this *SliceList) GetValue(index int) int {
	return this.slice[index]
}

func (this *SliceList) Print() {
	for index, val := range this.slice {
		fmt.Println("index:" + strconv.Itoa(index) + " " + "val:" + strconv.Itoa(val))
	}
}
