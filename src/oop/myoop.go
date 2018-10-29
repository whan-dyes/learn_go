package oop

import (
	. "fmt"
)

func MyOop() {
	Println(">>>>>>>>>>>>>>>>>>>>>>>>> MyOop <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	oop1()
	oop2()
	oop3()
	oop4()
	oop5()
}

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

func oop1() {
	Println("......... oop1 ..........")

	var a Integer = 1
	if a.Less(2) {
		Println("a.Less(2)=", a, "Less 2")
	}

	Println("a.Add(2)=", a.Add(2)) // a 值改变
	if a.Less(2) {
		Println("a.Less(2)=", a, "Less 2")
	} else {
		Println("a.Less(2)=", a, "not less 2")
	}

}

func oop2() {
	Println("......... oop2 ..........")

}

func oop3() {
	Println("......... oop3 ..........")

}

func oop4() {
	Println("......... oop4 ..........")

}

func oop5() {
	Println("......... oop5 ..........")

}
