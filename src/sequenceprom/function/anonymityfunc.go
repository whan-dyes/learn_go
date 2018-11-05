package function

import (
	. "fmt"
)

// abz:=func(a,b int,z float64) bool{
//   return (a*b)<int(z)
// }
func AnonymityFun() {
	Println("<<<<<<<<<<<<<<<<<<<<<AnonymityFun>>>>>>>>>>>>>>>>>>>>>>>>")
	anonymity1()
	anonymity2()
}

//--------------------------
func test() func() int {
	var x int //x无初始化，默认值为0
	return func() int {
		x++
		return x * x //退出闭包，匿名函数时，x不会释放，还会继续存在。
	}
}
func anonymity2() {
	Println("................anonymity2................")
	f := test()
	Println(f()) //1
	Println(f()) //4
	Println(f()) //9
	Println(f()) //16
	Println(f()) //25
}

func anonymity1() {
	Println("................anonymity1................")
	v1 := func(i int) int {
		return i * i
	}
	Println(v1(3))

	v2 := b(2)
	Println(v2(3))

	// for i := 0; i < 5; i++ {
	// 	defer func() { //defer 在函数返回后才运行，这时 i 已经变成了 5
	// 		Println("i=", i)
	// 	}()
	// }
	i1 := 0
	for ; i1 < 3; i1++ {
		func(i int) {
			Println("i1=", i)
		}(i1)
	}

	f := func(i int) int { //闭包以引用方式捕获外部变量，只要闭包在使用，变量就还会存在
		i1 += 100
		Println("i1=", i+i1)
		return i1
	}
	// f1(i1)

	f(i1)
}

func b(x int) func(int) int {
	return func(y int) int {
		return x + y
	}

}
