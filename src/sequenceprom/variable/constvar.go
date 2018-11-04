package variable

import . "fmt"

/*
常量的定义

常量的值在编译时就已经确定
常量的定义格式与变量基本相同
等号右侧必须是常量或者常量表达式
常量表达式中的函数必须是内置函数

常量的初始化规则与枚举

在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
使用相同的表达式不代表具有相同的值
iota是常量的计数器，从0开始，组中每定义1个常量自动递增1
通过初始化规则与iota可以达到枚举的效果
每遇到一个const关键字，iota就会重置为0
*/

// 在Go语言中，常量是指编译期间就已知且不可改变的值。常量可以是数值类型（包括整型、
// 浮点型和复数类型）、布尔类型、字符串类型等。
func ConstVar() {
	Println("---------------ConstVar-------------------")
	const Pi float64 = 3.14159265358979323846
	const zero = 0.0 // 无类型浮点常量
	const (
		size int64  = 1024
		eof         = -1 // 无类型整型常量
		Name string = "Joe"
		alen int    = len(Name) //常量可用内置函数初始化
	)

	Println("Name=", Name)

	const u, v float32 = 0, 3 // u = 0.0, v = 3.0，常量的多重赋值
	const a, b, c = 3, 4, "foo"
	Println("a=", a)
	/*a = 3, b = 4, c = "foo", 无类型整型和字符串常量
	Go的常量定义可以限定常量类型，但不是必需的。如果定义常量时没有指定类型，那么它
	与字面常量一样，是无类型常量。
	常量定义的右值也可以是一个在编译期运算的常量表达式，比如*/
	const mask = 1 << 3
	Println("mask=", mask)
	// 由于常量的赋值是一个编译期行为，所以右值不能出现任何需要运行期才能得出结果的表达
	// 式，比如试图以如下方式定义常量就会导致编译错误：
	//const Home = os.GetEnv("HOME")
	// 原因很简单，os.GetEnv()只有在运行期才能知道返回结果，在编译期并不能确定，所以
	// 无法作为常量定义的右值。

	// Go语言预定义了这些常量：true、false和iota。
	// iota比较特殊，可以被认为是一个可被编译器修改的常量，在每一个const关键字出现时被
	// 重置为0，然后在下一个const出现之前，每出现一次iota，其所代表的数字会自动增1。
	// 从以下的例子可以基本理解iota的用法：
	const ( // iota被重设为0
		c20 = iota // c0 == 0
		c21 = iota // c1 == 1
		c22 = iota // c2 == 2
	)
	const (
		a2 = 1 << iota // a == 1 (iota在每个const开头被重设为0)
		b2 = 1 << iota // b == 2
		c2 = 1 << iota // c == 4
	)
	const (
		u1         = iota * 42 // u == 0
		v1 float64 = iota * 42 // v == 42.0
		w1         = iota * 42 // w == 84
	)
	const x = iota // x == 0 (因为iota又被重设为0了)
	const y = iota // y == 0 (同上)
	// 如果两个const的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式。因此，上
	// 面的前两个const语句可简写为：
	const ( // iota被重设为0
		c10 = iota // c0 == 0
		c11        // c1 == 1
		c12        // c2 == 2
	)
	const (
		a1 = 1 << iota // a == 1 (iota在每个const开头被重设为0)
		b1             // b == 2
		c1             // c == 4
	)

	Println("b1=", b1)

	const (
		a71, b71 = 1, "2"
		c71, d   //等同 c71,d=1,"2"
	)

}
