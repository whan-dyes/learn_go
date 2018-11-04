package variable

import (
	"fmt"
)

/*
运算符优先级：
高到低
^   !     (一元运算符)
*   /  %  <<  >> &  &^   （二元运算符）
+  -  |  ^                （二元运算符）
==  !=  < <=  >=  >        （二元运算符）
<-                    (专用于channel)
&&              （短路运算符，第一个为 0 则返回 false，后面的不计算）
||               （短路运算符，第一个为 1 则返回 false，后面的不计算）

++ --  只能作为语句，而不能作为表达式，不能放在等号左右，要放在变量右边，左边不行

Go内置关键字（25个均为小写）

break        default           func        interface        select
case          defer              go           map               struct
chan          else                goto       package        switch
const         fallthrough    if             range             type
continue   for                  import    return             var


Go注释方法

// ：单行注释
/* */：多行注释


Go基本类型

    复数：complex64/complex128
    	- 长度：8/16字节
    足够保存指针的 32 位或 64 位整数型：uintptr

    其它值类型：
    	- array、struct、string
    引用类型：
    	- slice、map、chan

    接口类型：inteface
    函数类型：func

Go基本类型

    16位整型：int16/uint16
    	- 长度：2字节
    	- 取值范围：-32768~32767/0~65535
    32位整型：int32（rune）/uint32
    	- 长度：4字节
    	- 取值范围：-2^32/2~2^32/2-1/0~2^32-1
    64位整型：int64/uint64
    	- 长度：8字节
    	- 取值范围：-2^64/2~2^64/2-1/0~2^64-1
    浮点型：float32/float64
    	- 长度：4/8字节
    	- 小数位：精确到7/15小数位



Go基本类型

    布尔型：bool
    	- 长度：1字节
    	- 取值范围：true, false
    	- 注意事项：不可以用数字代表true或false

    整型：int/uint
    	- 根据运行平台可能为32或64位

    8位整型：int8/uint8
    	- 长度：1字节
    	- 取值范围：-128~127/0~255
    字节型：byte（uint8别名）


*/

func Operator() {
	var i, j int = 5, 3
	fmt.Println(i, " + ", j, "=", i+j)
	fmt.Println(i, " - ", j, "=", i-j)
	fmt.Println(i, " * ", j, "=", i*j)
	fmt.Println(i, " / ", j, "=", i/j)
	fmt.Println(i, " % ", j, "=", i%j)
	fmt.Println()
	//   ||,&&
	if i == 5 && j == 3 {
		fmt.Println(i == 3 && j == 2)
	}

	if i == 3 || j == 3 {
		fmt.Println(i == 3 || j == 2)
	}

	fmt.Println(i, " <", j, "=", i < j)
	fmt.Println(i, "< ", j, "=", i > j)
	fmt.Println(i, "== ", j, "=", i == j)
	fmt.Println(i, "<= ", j, "=", i <= j)
	fmt.Println(i, ">= ", j, "=", i >= j)
	fmt.Println(i, "!= ", j, "=", i != j)
	fmt.Println()

	/*
	   6:  0110
	   11: 1011
	   --------
	   &   0010
	   |   1111
	   ^   1101
	   &^  0100
	*/
	var x, y uint = 8, 1
	fmt.Println(x, " << ", y, "=", x<<y) //左移
	fmt.Println(x, " >> ", y, "=", x>>y)
	fmt.Println(x, "^ ", y, "=", x^y)    //异或  相同为 0
	fmt.Println(x, " &", y, "=", x&y)    //与   都是 1 则为 1
	fmt.Println(x, " | ", y, "=", x|y)   //或   都是 0 则为 0
	fmt.Println(x, " &^ ", y, "=", x&^y) //第二个数们为1时，第一个数对应位强制为0，为0时不变。
	var a uint8 = 255
	fmt.Println("^", a, "=", ^a) //取反
	fmt.Println()

}
