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

var x,y uint=8,1
fmt.Println(x," << ",y,"=",x<< y) //左移
fmt.Println(x," >> ",y,"=",x >>y)
fmt.Println(x,"^ ",y,"=",x ^ y)   //异或
fmt.Println(x," &",y,"=",x &y)    //与
fmt.Println(x," | ",y,"=",x| y)   //或
var a uint8=255
fmt.Println("^",a,"=",^a)         //取反
fmt.Println()


}
