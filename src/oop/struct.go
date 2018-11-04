package oop

import (
	. "fmt"
)

func MyStruct() {
	Println("---------------------MyStruct--------------------------")
	struct1()
	struct2()
	struct3()
	struct4()
	// struct5()
	// struct6()
	// struct7()
	// struct8()
	// struct9()
	// struct10()
}
func struct4() {
	Println("..............struct4  嵌入结构初始化..........")
	a := &teacher{Name: "joe", Age: 38, human: human{Sex: 0}}
	b := &student{Name: "Tom", Age: 18, human: human{Sex: 1}}
	a.Age = 41
	a.Sex = 1
	Println("teacher", a, " student", b)
}

func struct3() {
	Println("..............struct3  匿名结构..........")
	a := &struct {
		Name string
		Age  int
	}{Name: "Marry",
		Age: 35,
	}
	Println(a)
	//----------------------

	//包含匿名结构的结构成员初始化方法
	b := &person1{name: "Petty", age: 15}

	b.contact.Phone = "223344"
	b.contact.City = "BeiJing"

	Println(b)
}

func struct2() {
	Println("..............struct2  person..........")
	person1 := new(person)
	person2 := &person{}
	person3 := person{"John", 50}
	person4 := &person{Age: 23}
	person5 := NewPerson("Joe", 20)
	// s := person5.Area()
	// Println("person5.Area()=", s)
	Println("person1=", person1)
	Println("person2=", person2)

	Println("person3=", person3)
	person3.changePerson("Tom", 18)
	Println("person3=", person3)

	Println("person4=", person4)
	Println("person5=", person5)
}

func struct1() {
	Println("..............struct1 Rect..........")

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

/*
结构struct

Go 中的struct与C中的struct非常相似，并且Go没有class
使用 type <Name> struct{} 定义结构，名称遵循可见性规则
支持指向自身的指针类型成员
支持匿名结构，可用作成员或定义成员变量
匿名结构也可以用于map的值
可以使用字面值对结构进行初始化
允许直接通过指针来读写结构成员
相同类型的成员可进行直接拷贝赋值
支持 == 与 !=比较运算符，但不支持 > 或 <
支持匿名字段，本质上是定义了以某个类型名为名称的字段
嵌入结构作为匿名字段看起来像继承，但不是继承
可以使用匿名字段指针
*/
