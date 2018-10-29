package variable
import (
  . "fmt"
  "strconv"
  )

func StringVar(){
  Println("--------------------StringVar------------------------")
  var v2 string="hello"
  var v20="world"
  // v21：="Hello golang" invalid identifier character U+FF1A '：'
  v21 :="Hello golang"

  str := "Hello world" // 字符串也支持声明时进行初始化的做法
/*  str[0] = 'X' // 编译错误 str declared and not used,字符串的内容可以用类似于
数组下标的方式获取，但与数组不同，字符串的内容不能在初始化后被修改.*/

  Println("v2=",v2,"v20=",v20,"v21=",v21,"str=",str)

  var s1 string=v2 +" "+ v20
  Println("s1=",s1,"len(s1)=",len(s1),"s1[0]=",s1[0])
  // 字符串遍历
  for i:=0;i<len(s1);i++{
    ch:=s1[i]
    Println(i,ch)
  }
  Println()

  for i,ch:=range s1{
    Println(i,ch)
  }
  Println()

  var a int=65
  b:=strconv.Itoa(a)
  Println(b) //输出字串 65
  a,_=strconv.Atoi(b)
  Println("Atoi()=" ,a)

  var a1 int=65
  b1:=string(a1)
  Println(b1) //输出字串 A
}
