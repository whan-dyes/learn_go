package oop

type Rect struct {
	x, y          float64
	width, height float64
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

//------------------------------------------------------------------
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

//------------------------------------------------------------------

type MyInteger interface {
	Less(b Integer) bool
	Add(b Integer) Integer
}

//------------------------------------------------------------------
type person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *person {
	return &person{name, age}
}
func (per *person) changePerson(name string, age int) {
	per.Name = name
	per.Age = age

}

//-------------------------------------------------
type person1 struct {
	name    string
	age     int
	contact struct {
		Phone, City string
	}
}

//--------------------------------------------------
type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}

//---------------------------------------------------
type Test struct {
	str string "aaa"
}
