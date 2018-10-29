package variable
import (
  "fmt"
)

func Operator(){
  var i,j int=5,3
  fmt.Println(i," + ",j,"=",i+j)
  fmt.Println(i," - ",j,"=",i-j)
  fmt.Println(i," * ",j,"=",i*j)
  fmt.Println(i," / ",j,"=",i/j)
  fmt.Println(i," % ",j,"=",i%j)
  fmt.Println()
//   ||,&&
if i==5 && j==3 {
  fmt.Println(i==3 && j==2)
}

if i==3 || j==3 {
  fmt.Println(i==3 || j==2)
}

fmt.Println(i," <",j,"=",i < j)
fmt.Println(i,"< ",j,"=",i> j)
fmt.Println(i,"== ",j,"=",i== j)
fmt.Println(i,"<= ",j,"=",i<= j)
fmt.Println(i,">= ",j,"=",i >=j)
fmt.Println(i,"!= ",j,"=",i!= j)
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
var x,y uint=8,1
fmt.Println(x," << ",y,"=",x<< y) //左移
fmt.Println(x," >> ",y,"=",x >>y)
fmt.Println(x,"^ ",y,"=",x ^ y)   //异或  相同为 0
fmt.Println(x," &",y,"=",x &y)    //与   都是 1 则为 1
fmt.Println(x," | ",y,"=",x| y)   //或   都是 0 则为 0
fmt.Println(x," &^ ",y,"=",x&^ y)   //第二个数们为1时，第一个数对应位强制为0，为0时不变。
var a uint8=255
fmt.Println("^",a,"=",^a)         //取反
fmt.Println()


}

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
*/
