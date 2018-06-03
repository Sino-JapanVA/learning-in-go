package List

import (
	"fmt"
	"strconv"
)

type Node struct {
	Val  int
	prev *Node
	next *Node
}

type NodeList struct {
	head     *Node
	fence    *Node
	tail     *Node
	listSize int
}

func NewNodeList() *NodeList {
	return &NodeList{nil, nil, nil, 0}
}

func (this *NodeList) Append(Val int) {
	if this.head == nil {
		this.head = &Node{Val, nil, nil}
		this.fence = this.head
		this.tail = this.head
		this.listSize++
	} else {
		node := &Node{Val, this.tail, nil}
		this.tail.next = node
		this.tail = node
		this.fence = this.tail
		this.listSize++
	}
}

func (this *NodeList) Size() int {
	return this.listSize
}

func (this *NodeList) RemoveAll() {
	this.head = nil
	this.tail = nil
	this.fence = nil
}

func (this *NodeList) Print() {
	current := this.head
	for current.next != nil {
		fmt.Println("Value:" + strconv.Itoa(current.Val))
		current = current.next
	}
}

func (this *NodeList) GetVal(ele *int) bool {
	fmt.Println(ele)
	*ele = this.fence.Val
	fmt.Println(*ele)
	return true
}

func (this *NodeList) Find(ele **Node, Val int) bool {
	current := this.head
	for current.next != nil {
		if current.Val == Val {
			// 让这个指向指针的指针指向current的地址
			*ele = current
			return true
		}
		current = current.next
	}
	return false
}
