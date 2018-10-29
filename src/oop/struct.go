package oop

import (
	"fmt"
)

type INT int //类型别名

type A struct {
	B
	Name string
}

type B struct {
	Name string
}

func MyStruct() {
	a := &A{B: B{Name: "A_B Name"}, Name: "Joe A"}

	c := struct {
		Age int
	}{
		Age: 28,
	}
	fmt.Println("c=", c.Age)

	a.Print()
	fmt.Println(a.Name, ",", a.B.Name)

	b := &B{Name: "John B"}
	b.Print()
	fmt.Println(b.Name)

	var i INT
	i = 100
	i.Print()
}

func (a *A) Print() { //引用传递
	a.Name = "joe AA"
	fmt.Println("A=", a)
}

func (b B) Print() { //值传递
	b.Name = "John BB"
	fmt.Println("B=", b)
}

func (i INT) Print() { //值传递
	fmt.Println("INT=", i)
}
