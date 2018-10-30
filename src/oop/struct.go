package oop

import (
	. "fmt"
)

func MyStruct() {
	Println("---------------------MyStruct--------------------------")
	struct1()
	// struct2()
	// struct3()
	// struct4()
	// struct5()
	// struct6()
	// struct7()
	// struct8()
	// struct9()
	// struct10()
}

func struct1() {
	Println("..............struct1..........")
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
	rect5 := NewRect(0, 0, 10, 20)
	s := rect5.Area()
	Println("rect5.Area()=", s)
	Println("rect1=", rect1)
	Println("rect2=", rect2)
	Println("rect3=", rect3)
	Println("rect4=", rect4)
	Println("rect5=", rect5)
}
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}
