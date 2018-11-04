package variable

import (
	. "fmt"
)

/*
指针

        Go虽然保留了指针，但与其它编程语言不同的是，在Go当中不
支持指针运算以及”->”运算符，而直接采用”.”选择符来操作指针
目标对象的成员

操作符”&”取变量地址，使用”*”通过指针间接访问目标对象
默认值为 nil 而非 NULL


递增递减语句

        在Go当中，++ 与 -- 是作为语句而并不是作为表达式
*/

func Pointer() {
	Println("-------------------pointer------------------")
	test1()
}

func test1() {
	a := 1
	var p *int = &a
	Println("p=", p, "     *p=", *p)
}
func test2() {

}
func test3() {

}
func test4() {

}
func test5() {

}
func test6() {

}
func test7() {

}

/*
指针默认值为 nil 而非 NULL
操作符 & 取变量地址，使用 * 通过指针间接访问目标对象
*/
