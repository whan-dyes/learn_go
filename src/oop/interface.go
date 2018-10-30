package oop

import (
	. "fmt"
)

func MyInterface() {
	Println("---------------------Myinterface--------------------------")
	interface1()
	// interface2()
	// interface3()
	// interface4()
	// interface5()
	// interface6()
	// interface7()
	// interface8()
	// interface9()
	// interface10()
}

func interface1() {
	Println("..............interface1..........")
	var a Integer = 0
	var i MyInteger = &a
	Println(i)
	s := i.Less(2)
	Println(s)

}

func interface2() {
	Println("..............interface2..........")
}
func interface3() {
	Println("..............interface3..........")
}
