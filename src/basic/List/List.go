package List

type List struct {
	listSize  int
	maxSize   int
	listArray []int
	fence     int
}

func (this *List) New(size int) *List {
	this.listSize = size
	return this
}

func (this *List) Clear() bool {
	this.listArray = make([]int, this.maxSize)
	return true
}

func (this *List) Insert(val int) bool {
	return true
}

func (this *List) Append(val int) bool {
	return true
}

func (this *List) Remove() bool {
	return true
}

func (this *List) SetStart() bool {
	return true
}

func (this *List) SetEnd() bool {
	return true
}

func (this *List) Prev() {

}

func (this *List) Next() {

}

func (this *List) LeftLength() int {
	return 0
}

func (this *List) RightLength() int {
	return 0
}

func (this *List) SetPos(pos int) bool {
	if pos >= 0 && pos <= this.listSize {
		this.fence = pos
	}
	return (pos >= 0) && (pos <= this.listSize)
}

func (this *List) GetValue(val *int) bool {
	if this.RightLength() == 0 {
		return false
	} else {
		val = &this.listArray[this.fence]
		return true
	}
}

func (this *List) Print() {

}

func (this *List) Find() bool {
	return true
}
