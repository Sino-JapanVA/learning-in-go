package Stack

import (
	"basic/List"
)

type Stack struct {
	size int
	top  int
	list *List.SliceList
}

func NewStack(size int) *Stack {
	list := List.NewSliceList()
	return &Stack{size, 0, list}
}

func (this *Stack) Clear() {
	this.top = 0
	this.list.Clear()
}

func (this *Stack) Push(ele int) {
	this.list.Append(ele)
	this.top++
}

func (this *Stack) Pop(ele *int) bool {
	if ok := this.list.Remove(ele, this.top-1); ok {
		this.top--
		return true
	} else {
		return false
	}
}

func (this *Stack) GetValue() int {
	return this.list.GetValue(this.top - 1)
}

func (this *Stack) Print() {
	this.list.Print()
}
