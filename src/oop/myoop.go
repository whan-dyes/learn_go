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

/*
方法method

Go 中虽没有class，但依旧有method
通过显示说明receiver来实现与某个类型的组合
只能为同一个包中的类型定义方法
Receiver 可以是类型的值或者指针
不存在方法重载
可以使用值或指针来调用方法，编译器会自动完成转换
从某种意义上来说，方法是函数的语法糖，因为receiver其实就是
方法所接收的第1个参数（Method Value vs. Method Expression）
如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
类型别名不会拥有底层类型所附带的方法
方法可以调用结构中的非公开字段
*/
