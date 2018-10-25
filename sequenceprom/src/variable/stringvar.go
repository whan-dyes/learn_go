package variable
import "fmt"

func StringVar(){
  var v2 string="hello"
  var v20="world"
  // v21：="Hello golang" invalid identifier character U+FF1A '：'
  v21 :="Hello golang"

  str := "Hello world" // 字符串也支持声明时进行初始化的做法
  // str[0] = 'X' // 编译错误 str declared and not used,字符串的内容可以用类似于数组下标的方式获取，但与数组不同，字符串的内容不能在初始
  // 化后被修改.

  fmt.Println("v2=",v2,"v20=",v20,"v21=",v21,"str=",str)

  var s1 string=v2 +" "+ v20
  fmt.Println("s1=",s1,"len(s1)=",len(s1),"s1[0]=",s1[0])
  // 字符串遍历
  for i:=0;i<len(s1);i++{
    ch:=s1[i]
    fmt.Println(i,ch)
  }
  fmt.Println()

  for i,ch:=range s1{
    fmt.Println(i,ch)
  }
  fmt.Println()
}
