package function

import (
	. "fmt"
)

func MyFunc() {
	Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>> MyFunc <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	myfunc1()
	myfunc2()
	s1 := []int{1, 2, 3}
	myfunc3(s1, 1, 2, 3)

	x := 11
	myfunc4(&x)
	myfunc5()
	myfunc6(10, 20, add) //参数为函数名，实现多态
	myfunc6(20, 10, minus)
	// s := []int{1, 2, 3}
	// myfunc3(s, 4, 5, 6, 7, 8)
	// Println(s)
	//
	// a := 1
	// Println("Myfunc() a=", a)
	// M := myfunc4 //golang中函数也是一种类型
	// M(&a)
	// Println("Myfunc() a=", a)
}

func myfunc1() (a, b, c int) {
	Println("............ myfunc1() ...........")
	a, b, c = 1, 2, 3
	return
}

func myfunc2() (int, int, int) {
	Println("............ myfunc2() ...........")
	a, b, c := 1, 2, 3
	return a, b, c
}

func myfunc3(b []int, a ...int) []int { // ...int 相当于 slice，但是值传递，不是引用
	Println("............ myfunc3() ...........")
	Println(b)
	b[0] = 100

	Println(a)
	a[0] = 100

	return a
}

func myfunc4(a *int) {
	Println("............ myfunc4() ...........")
	*a = 100
	Println("*a=", *a)
}

//------------函数类型示例--------------
func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}

type FuncType func(int, int) int

func myfunc5() {
	Println("............ myfunc5() 函数类型...........")
	var f FuncType
	f = add
	Println(f(20, 10))

	f = minus
	Println(f(20, 10))
}
func myfunc6(a, b int, f FuncType) {
	Println("............ myfunc6() 回调函数，传入函数名参数..........")

	Println("f=", f(a, b))
}

/*
函数function

Go 函数 不支持 嵌套、重载和默认参数
但支持以下特性：

        无需声明原型、不定长度变参、多返回值、命名返回值参数
        匿名函数、闭包

定义函数使用关键字 func，且左大括号不能另起一行
函数也可以作为一种类型使用
*/
