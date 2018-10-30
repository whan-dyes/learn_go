package oop

type Rect struct {
	x, y          float64
	width, height float64
}

//-----------------
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

// 为Go语言和C语言一样，类型都是基于值传递的。要想修改变量的值，只能
// 传递指针。
func (a *Integer) Add(b Integer) Integer {
	*a += b
	return *a
}

//-------------------
type MyInteger interface {
	Less(b Integer) bool
	Add(b Integer) Integer
}
