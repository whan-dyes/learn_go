package main

import (
	"channel"
	. "fmt" //省略包名直接调用
	"function"
	myoop "oop" //包别名
	"reflection"
	"testinterface"
	"variable"
)

func main() {
	Println("Integervar..................................................")
	variable.IntegerVar()
	Println()

	Println("Stringvar..................................................")
	variable.StringVar()
	Println()

	Println("Anonymityvar..................................................")
	variable.AnonymityVar()
	Println()

	Println("Constvar..................................................")
	variable.ConstVar()
	Println()

	Println("EnumVar..................................................")
	variable.EnumVar()
	// fmt.Println("Monday=",variable.Monday)
	Println()

	Println("BoolVar..................................................")
	variable.BoolVar()
	Println()
	// var v3 [10]int // 数组
	// var v4 []int // 数组切片
	// var v5 struct {
	// f int
	// }
	// var v6 *int // 指针
	// var v7 map[string]int // map，key为string类型，value为int类型
	// var v8 func(a int) int
	// v8 = func (a int)int{
	//   fmt.Println("a=",a)
	//   return a
	// }
	// fmt.Println("v8 return=",v8(100),"v8 func=",v8)

	Println("floatVar..................................................")
	variable.FloatVar()
	Println()

	function.A()

	Println("operator...................................................")
	variable.Operator()
	Println()

	Println("ComplexVar...................................................")
	variable.ComplexVar()
	Println()

	Println("ArrayVar...................................................")
	variable.ArrayVar()
	Println()

	Println("Slice...................................................")
	variable.Slice()
	Println()

	Println("MyMap............................................")
	variable.MyMap()
	Println()

	Println("MySwitch................................................")
	function.MySwitch()
	Println()

	Println("OopTest................................................")
	myoop.OopTest()
	Println()

	Println("Anonymityfunc............................................")
	function.A()
	Println()

	Println("MyStruct................................................")
	myoop.MyStruct()
	Println()

	Println("MyInterface................................................")
	testinterface.MyInterface()
	Println()

	Println("MyReflection................................................")
	reflection.MyReflection()
	Println()

	Println("MyChannel................................................")
	channel.MyChannel()
	Println()
}
