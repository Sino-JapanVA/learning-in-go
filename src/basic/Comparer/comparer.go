package Comparer

type Comparer interface {
	It(x, y int) bool
	Eq(x, y int) bool
	Gt(x, y int) bool
}
