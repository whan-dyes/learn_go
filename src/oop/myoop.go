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
