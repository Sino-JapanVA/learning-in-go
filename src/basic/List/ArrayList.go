package List

import (
	"fmt"
)

type ArrayList struct {
	listSize  int
	maxSize   int
	listArray []int
	fence     int
}

func NewArrayList(size int) *ArrayList {
	list := &ArrayList{0, size, make([]int, size), 0}
	return list
}

func (this *ArrayList) Clear() bool {
	this.listArray = make([]int, this.maxSize)
	return true
}

func (this *ArrayList) Insert(val int) bool {
	return true
}

func (this *ArrayList) Append(val int) bool {
	if this.listSize < this.maxSize {
		this.listArray = append(this.listArray, val)
		return true
	} else {
		return false
	}
}

func (this *ArrayList) Remove() bool {
	return true
}

func (this *ArrayList) SetStart() bool {
	return true
}

func (this *ArrayList) SetEnd() bool {
	return true
}

func (this *ArrayList) Prev() {

}

func (this *ArrayList) Next() {

}

func (this *ArrayList) LeftLength() int {
	return 0
}

func (this *ArrayList) RightLength() int {
	return 0
}

func (this *ArrayList) SetPos(pos int) bool {
	if pos >= 0 && pos <= this.listSize {
		this.fence = pos
	}

	return (pos >= 0) && (pos <= this.listSize)
}

func (this *ArrayList) GetValue(val *int) bool {
	if this.RightLength() == 0 {
		return false
	} else {
		val = &this.listArray[this.fence]
		return true
	}
}

func (this *ArrayList) Print() {
	for index, val := range this.listArray {
		fmt.Println(index + ':' + val)
	}
}

func (this *ArrayList) Find() bool {
	return true
}
